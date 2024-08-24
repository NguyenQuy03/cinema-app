package business

import (
	"context"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/modules/user/model"
	"github.com/NguyenQuy03/cinema-app/server/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginStorage interface {
	GetUser(ctx context.Context, conds map[string]interface{}) (*model.User, error)
}

type loginUserBiz struct {
	storage UserLoginStorage
}

func NewLoginUserBiz(storage UserLoginStorage) *loginUserBiz {
	return &loginUserBiz{
		storage: storage,
	}
}

func (biz *loginUserBiz) AuthenticateUser(ctx context.Context, data *model.UserLogin, authResponse *model.AuthResponse) error {

	email := strings.TrimSpace(data.Email)
	password := strings.TrimSpace(data.Password)

	if email == "" || password == "" {
		return model.ErrEmailOrPassMissing
	}

	if !utils.IsValidEmail(email) {
		return model.ErrEmailInvalid
	}

	// Find User by email
	user, err := biz.storage.GetUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return model.ErrLoginFailure
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.ErrLoginFailure
	}

	accessToken, err := utils.GenerateAccessToken(user)

	if err != nil {
		return model.ErrGenerateToken
	}

	authResponse.AccessToken = accessToken

	refreshToken, err := utils.GenerateRefreshToken(user)

	authResponse.RefreshToken = refreshToken

	if err != nil {
		return model.ErrGenerateToken
	}

	return nil
}
