package business

import (
	"context"
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type UpdateSessionStorage interface {
	GetUserSession(ctx context.Context, email string) (map[string]string, error)
	StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration int) error
}

type JWTHandler interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
	GenerateAccessToken(sub string) (string, int, error)
	GenerateRefreshToken(sub string) (string, int, error)
	ParseToken(tokenString string) (claims *jwt.RegisteredClaims, err error)
	CompareToken(token1, token2 string) (bool, error)
}

type refreshTokenBiz struct {
	getUserStorage       GetUserStorage
	updateSessionStorage UpdateSessionStorage
	jwtHandler           JWTHandler
}

func NewRefreshTokenBiz(getUserStorage GetUserStorage, updateSessionStorage UpdateSessionStorage, jwtHandler JWTHandler) *refreshTokenBiz {
	return &refreshTokenBiz{
		getUserStorage:       getUserStorage,
		updateSessionStorage: updateSessionStorage,
		jwtHandler:           jwtHandler,
	}
}

func (biz *refreshTokenBiz) RefreshToken(c context.Context, req *http.Request, refreshToken string) (*model.AuthResponse, error) {
	// Validate request token
	token, err := biz.jwtHandler.ValidateToken(refreshToken)

	if err != nil || token == nil {
		return nil, err
	}

	// Extract email from token
	claims, err := biz.jwtHandler.ParseToken(refreshToken)
	email := claims.Subject

	if err != nil {
		return nil, err
	}

	// Get prev refresh token in redis
	userSession, err := biz.updateSessionStorage.GetUserSession(c, claims.Subject)

	if err != nil {
		return nil, err
	}

	prevRefreshToken := userSession[common.RefreshToken]

	// Compare Tokens
	_, err = biz.jwtHandler.CompareToken(prevRefreshToken, refreshToken)

	if err != nil {
		return nil, err
	}

	// Generate and return new access token
	user, err := biz.getUserStorage.GetUser(c, map[string]interface{}{
		"email": email,
	})

	if err != nil {
		return nil, err
	}

	accessToken, expAcTokenSecs, err := biz.jwtHandler.GenerateAccessToken(user.Email)

	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken: model.Token{
			Token:     accessToken,
			ExpiredIn: expAcTokenSecs,
		},
	}, nil
}
