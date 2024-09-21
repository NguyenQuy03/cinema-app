package business

import (
	"context"
	"strings"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"golang.org/x/crypto/bcrypt"
)

type StoreSessionStorage interface {
	StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration time.Duration) error
}

type loginUserBiz struct {
	userStorage         GetUserStorage
	storeSessionStorage StoreSessionStorage
}

func NewLoginUserBiz(userStorage GetUserStorage, storeSessionStorage StoreSessionStorage) *loginUserBiz {
	return &loginUserBiz{
		userStorage:         userStorage,
		storeSessionStorage: storeSessionStorage,
	}
}

func (biz *loginUserBiz) AuthenticateUser(ctx context.Context, data *model.UserLogin, authResponse *model.AuthResponse) error {

	email := strings.TrimSpace(data.Email)
	password := data.Password

	if email == "" || password == "" {
		return model.ErrEmailOrPassMissing
	}

	// Find User by email
	user, err := biz.userStorage.GetUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return model.ErrLoginFailure
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.ErrLoginFailure
	}

	// Genarate Tokens
	handleTokenBiz := HandleTokenBiz{}
	accessToken, err := handleTokenBiz.GenerateAccessToken(user)

	if err != nil {
		return model.ErrGenerateToken
	}

	authResponse.AccessToken = accessToken

	refreshToken, err := handleTokenBiz.GenerateRefreshToken(user)

	authResponse.RefreshToken = refreshToken

	if err != nil {
		return model.ErrGenerateToken
	}

	// Use Redis to store session
	err = biz.storeSessionStorage.StoreUserSession(
		ctx,
		user.Email,
		map[string]interface{}{
			model.RefreshToken: refreshToken,
		},
		time.Until(model.RefreshTokenMaxAge),
	)

	if err != nil {
		return err
	}

	return nil
}
