package business

import (
	"context"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"golang.org/x/crypto/bcrypt"
)

type StoreSessionStorage interface {
	StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration int) error
}

type loginUserBiz struct {
	userStorage         GetUserStorage
	storeSessionStorage StoreSessionStorage
	jwtProvider         JWTProvider
}

func NewLoginUserBiz(userStorage GetUserStorage, storeSessionStorage StoreSessionStorage, jwtProvider JWTProvider) *loginUserBiz {
	return &loginUserBiz{
		userStorage:         userStorage,
		storeSessionStorage: storeSessionStorage,
		jwtProvider:         jwtProvider,
	}
}

func (biz *loginUserBiz) Login(ctx context.Context, data *model.UserLogin) (*model.AuthResponse, error) {

	email := strings.TrimSpace(data.Email)
	password := data.Password

	if email == "" || password == "" {
		return nil, model.ErrEmailOrPassMissing
	}

	// Find User by email
	user, err := biz.userStorage.GetUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, model.ErrLoginFailure
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, model.ErrLoginFailure
	}

	// Genarate Tokens
	accessToken, expAcTokenSecs, err := biz.jwtProvider.GenerateAccessToken(user.Email)

	if err != nil {
		return nil, err
	}

	refreshToken, expReTokenSecs, err := biz.jwtProvider.GenerateRefreshToken(user.Email)

	if err != nil {
		return nil, err
	}

	// Use Redis to store session
	err = biz.storeSessionStorage.StoreUserSession(
		ctx,
		user.Email,
		map[string]interface{}{
			common.RefreshToken: refreshToken,
		},
		expReTokenSecs,
	)

	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken: model.Token{
			Token:     accessToken,
			ExpiredIn: expAcTokenSecs,
		},
		RefreshToken: model.Token{
			Token:     refreshToken,
			ExpiredIn: expReTokenSecs,
		},
	}, nil
}
