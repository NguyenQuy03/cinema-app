package business

import (
	"context"
	"testing"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mocking dependencies
type mockUserStorage struct {
	mock.Mock
}

func (m *mockUserStorage) GetUser(ctx context.Context, filters map[string]interface{}) (*model.User, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).(*model.User), args.Error(1)
}

type mockStoreSessionStorage struct {
	mock.Mock
}

func (m *mockStoreSessionStorage) StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration int) error {
	return m.Called(ctx, key, infors, expiration).Error(0)
}

type mockJWTProvider struct {
	mock.Mock
}

// ParseToken implements JWTProvider.
func (m *mockJWTProvider) ParseToken(tokenString string) (claims *common.CustomClaims, err error) {
	panic("unimplemented")
}

// ValidateToken implements JWTProvider.
func (m *mockJWTProvider) ValidateToken(tokenString string) (*jwt.Token, error) {
	panic("unimplemented")
}

func (m *mockJWTProvider) GenerateAccessToken(email string, isAdmin bool) (string, int, error) {
	args := m.Called(email, isAdmin)
	return args.String(0), args.Int(1), args.Error(2)
}

func (m *mockJWTProvider) GenerateRefreshToken(email string, isAdmin bool) (string, int, error) {
	args := m.Called(email, isAdmin)
	return args.String(0), args.Int(1), args.Error(2)
}

func TestLoginUserBiz_Login(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Test case: Successful login
	t.Run("successful login", func(t *testing.T) {
		mockUserStorage := new(mockUserStorage)
		mockStoreSessionStorage := new(mockStoreSessionStorage)
		mockJWTProvider := new(mockJWTProvider)

		biz := business.NewLoginUserBiz(mockUserStorage, mockStoreSessionStorage, mockJWTProvider)

		password := "123123"
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{Email: "test@example.com", Password: string(hashedPassword), RoleCode: model.UserRoleCode}

		mockUserStorage.On("GetUser", ctx, map[string]interface{}{"email": user.Email}).Return(user, nil)
		mockJWTProvider.On("GenerateAccessToken", user.Email, false).Return("accessToken", 3600, nil)
		mockJWTProvider.On("GenerateRefreshToken", user.Email, false).Return("refreshToken", 7200, nil)
		mockStoreSessionStorage.On("StoreUserSession", ctx, user.Email, mock.Anything, 7200).Return(nil)

		data := &model.UserLogin{Email: "test@example.com", Password: password}
		resp, err := biz.Login(ctx, data)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "accessToken", resp.AccessToken.Token)
		assert.Equal(t, "refreshToken", resp.RefreshToken.Token)
	})

	// Test case: Missing email
	t.Run("missing email", func(t *testing.T) {
		mockUserStorage := new(mockUserStorage)
		mockStoreSessionStorage := new(mockStoreSessionStorage)
		mockJWTProvider := new(mockJWTProvider)

		biz := business.NewLoginUserBiz(mockUserStorage, mockStoreSessionStorage, mockJWTProvider)

		data := &model.UserLogin{Email: "", Password: "123123"}
		resp, err := biz.Login(ctx, data)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// Test case: Missing password
	t.Run("missing password", func(t *testing.T) {
		mockUserStorage := new(mockUserStorage)
		mockStoreSessionStorage := new(mockStoreSessionStorage)
		mockJWTProvider := new(mockJWTProvider)

		biz := business.NewLoginUserBiz(mockUserStorage, mockStoreSessionStorage, mockJWTProvider)

		data := &model.UserLogin{Email: "test@example.com", Password: ""}
		resp, err := biz.Login(ctx, data)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// Test case: Invalid password
	t.Run("invalid password", func(t *testing.T) {
		mockUserStorage := new(mockUserStorage)
		mockStoreSessionStorage := new(mockStoreSessionStorage)
		mockJWTProvider := new(mockJWTProvider)

		biz := business.NewLoginUserBiz(mockUserStorage, mockStoreSessionStorage, mockJWTProvider)

		mockUserStorage.On("GetUser", ctx, map[string]interface{}{"email": "test@example.com"}).Return(&model.User{Email: "test@example.com", Password: "hashedPassword"}, nil)

		data := &model.UserLogin{Email: "test@example.com", Password: "wrongPassword"}
		resp, err := biz.Login(context.Background(), data)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
