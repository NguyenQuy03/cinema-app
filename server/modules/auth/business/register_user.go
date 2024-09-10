package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/utils/mailUtil"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserStorage interface {
	GetUser(ctx context.Context, conds map[string]interface{}) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserRegister) error
}

type registerUserBiz struct {
	storage RegisterUserStorage
}

func NewRegisterUserBiz(storage RegisterUserStorage) *registerUserBiz {
	return &registerUserBiz{
		storage: storage,
	}
}

func (biz *registerUserBiz) RegisterUser(ctx context.Context, data *model.UserRegister) error {
	email := data.Email
	password := data.Password

	if !mailUtil.IsValidEmail(email) {
		return model.ErrEmailInvalid
	}

	if len(password) < 6 {
		return model.ErrShortPass
	}

	// Check user has already exist in DB or not
	_, err := biz.storage.GetUser(ctx, map[string]interface{}{"email": email})
	if err == nil {
		return model.ErrUserExisted
	}

	// Hash password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return model.ErrHashPassword
	}

	data.Password = string(hashPass)

	// Register new user
	if err := biz.storage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.UserEntityName)
	}

	return nil
}
