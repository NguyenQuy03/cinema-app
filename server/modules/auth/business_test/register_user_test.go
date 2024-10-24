package business

import (
	"context"
	"errors"
	"testing"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRegisterUserStorage struct {
	mock.Mock
}

func (m *MockRegisterUserStorage) GetUser(ctx context.Context, conds map[string]interface{}) (*model.User, error) {
	args := m.Called(ctx, conds)
	if args.Get(0) != nil {
		return args.Get(0).(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRegisterUserStorage) CreateUser(ctx context.Context, data *model.UserRegister) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestRegisterUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("successful registration", func(t *testing.T) {
		storage := new(MockRegisterUserStorage)
		biz := business.NewRegisterUserBiz(storage)

		data := &model.UserRegister{
			Email:    "test@example.com",
			Password: "123123",
		}

		storage.On("GetUser", ctx, map[string]interface{}{"email": data.Email}).Return(nil, common.ErrRecordNotFound)
		storage.On("CreateUser", ctx, data).Return(nil)

		err := biz.RegisterUser(ctx, data)

		require.NoError(t, err)
		storage.AssertCalled(t, "CreateUser", ctx, data)
	})

	t.Run("invalid email format", func(t *testing.T) {
		storage := new(MockRegisterUserStorage)
		biz := business.NewRegisterUserBiz(storage)

		data := &model.UserRegister{
			Email:    "invalid-email",
			Password: "password123",
		}

		err := biz.RegisterUser(ctx, data)

		require.ErrorIs(t, err, model.ErrEmailInvalid)
	})

	t.Run("password too short", func(t *testing.T) {
		storage := new(MockRegisterUserStorage)
		biz := business.NewRegisterUserBiz(storage)

		data := &model.UserRegister{
			Email:    "test@example.com",
			Password: "short",
		}

		err := biz.RegisterUser(ctx, data)

		require.ErrorIs(t, err, model.ErrShortPass)
	})

	t.Run("user already exists", func(t *testing.T) {
		storage := new(MockRegisterUserStorage)
		biz := business.NewRegisterUserBiz(storage)

		data := &model.UserRegister{
			Email:    "test@example.com",
			Password: "password123",
		}

		existingUser := &model.User{
			Email: data.Email,
		}

		storage.On("GetUser", ctx, map[string]interface{}{"email": data.Email}).Return(existingUser, nil)

		err := biz.RegisterUser(ctx, data)

		require.ErrorIs(t, err, model.ErrUserExisted)
	})

	t.Run("storage error when creating user", func(t *testing.T) {
		storage := new(MockRegisterUserStorage)
		biz := business.NewRegisterUserBiz(storage)

		data := &model.UserRegister{
			Email:    "test@example.com",
			Password: "password123",
		}

		storage.On("GetUser", ctx, map[string]interface{}{"email": data.Email}).Return(nil, common.ErrRecordNotFound)
		storage.On("CreateUser", ctx, data).Return(errors.New("db error"))

		err := biz.RegisterUser(ctx, data)

		require.Error(t, err)
		require.ErrorContains(t, err, "db error")
	})
}
