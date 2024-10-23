package business

import (
	"context"
	"net/mail"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
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
	email := strings.TrimSpace(data.Email)
	password := data.Password

	// Validate email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return model.ErrEmailInvalid
	}

	// Validate password
	if len(password) < 6 {
		return model.ErrShortPass
	}

	// Check user has already exist in DB or not
	_, err = biz.storage.GetUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil && err != common.ErrRecordNotFound {
		return model.ErrUserExisted
	}

	// Hash password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

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
