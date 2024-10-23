package business

import (
	"context"
	"testing"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mocking dependencies
type MockUserStorage struct {
	mock.Mock
}

func (m *MockUserStorage) GetUser(ctx context.Context, filters map[string]interface{}) (*model.User, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).(*model.User), args.Error(1)
}

type MockStoreSessionStorage struct {
	mock.Mock
}

func (m *MockStoreSessionStorage) StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration int) error {
	return m.Called(ctx, key, infors, expiration).Error(0)
}

type MockJWTProvider struct {
	mock.Mock
}

// CompareToken implements JWTProvider.
func (m *MockJWTProvider) CompareToken(token1 string, token2 string) (bool, error) {
	panic("unimplemented")
}

// ParseToken implements JWTProvider.
func (m *MockJWTProvider) ParseToken(tokenString string) (claims *common.CustomClaims, err error) {
	panic("unimplemented")
}

// ValidateToken implements JWTProvider.
func (m *MockJWTProvider) ValidateToken(tokenString string) (*jwt.Token, error) {
	panic("unimplemented")
}

func (m *MockJWTProvider) GenerateAccessToken(email string, isAdmin bool) (string, int, error) {
	args := m.Called(email, isAdmin)
	return args.String(0), args.Int(1), args.Error(2)
}

func (m *MockJWTProvider) GenerateRefreshToken(email string, isAdmin bool) (string, int, error) {
	args := m.Called(email, isAdmin)
	return args.String(0), args.Int(1), args.Error(2)
}

func TestLoginUserBiz_Login(t *testing.T) {
	mockUserStorage := new(MockUserStorage)
	mockStoreSessionStorage := new(MockStoreSessionStorage)
	mockJWTProvider := new(MockJWTProvider)

	biz := NewLoginUserBiz(mockUserStorage, mockStoreSessionStorage, mockJWTProvider)

	// Test case: Successful login
	password := "123123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &model.User{Email: "quy@gmail.com", Password: string(hashedPassword), RoleCode: model.UserRoleCode}

	mockUserStorage.On("GetUser", mock.Anything, mock.Anything).Return(user, nil)
	mockJWTProvider.On("GenerateAccessToken", user.Email, false).Return("accessToken", 3600, nil)
	mockJWTProvider.On("GenerateRefreshToken", user.Email, false).Return("refreshToken", 7200, nil)
	mockStoreSessionStorage.On("StoreUserSession", mock.Anything, user.Email, mock.Anything, 7200).Return(nil)

	data := &model.UserLogin{Email: "test@example.com", Password: password}
	resp, err := biz.Login(context.Background(), data)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "accessToken", resp.AccessToken.Token)
	assert.Equal(t, "refreshToken", resp.RefreshToken.Token)

	// Test case: Missing email
	data = &model.UserLogin{Email: "", Password: password}
	resp, err = biz.Login(context.Background(), data)

	assert.Error(t, err)
	assert.Nil(t, resp)

	// Test case: Missing password
	data = &model.UserLogin{Email: "test@example.com", Password: ""}
	resp, err = biz.Login(context.Background(), data)

	assert.Error(t, err)
	assert.Nil(t, resp)

	// Test case: Invalid password
	mockUserStorage.On("GetUser", mock.Anything, mock.Anything).Return(user, nil)
	mockJWTProvider.On("GenerateAccessToken", user.Email, false).Return("accessToken", 3600, nil)
	mockJWTProvider.On("GenerateRefreshToken", user.Email, false).Return("refreshToken", 7200, nil)
	mockStoreSessionStorage.On("StoreUserSession", mock.Anything, user.Email, mock.Anything, 7200).Return(nil)

	data = &model.UserLogin{Email: "test@example.com", Password: "wrongPassword"}
	resp, err = biz.Login(context.Background(), data)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
