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

type JWTProvider interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
	GenerateAccessToken(sub string, isAdmin bool) (string, int, error)
	GenerateRefreshToken(sub string, isAdmin bool) (string, int, error)
	ParseToken(tokenString string) (claims *common.CustomClaims, err error)
	CompareToken(token1, token2 string) (bool, error)
}

type refreshTokenBiz struct {
	getUserStorage       GetUserStorage
	updateSessionStorage UpdateSessionStorage
	jwtProvider          JWTProvider
}

func NewRefreshTokenBiz(getUserStorage GetUserStorage, updateSessionStorage UpdateSessionStorage, jwtProvider JWTProvider) *refreshTokenBiz {
	return &refreshTokenBiz{
		getUserStorage:       getUserStorage,
		updateSessionStorage: updateSessionStorage,
		jwtProvider:          jwtProvider,
	}
}

func (biz *refreshTokenBiz) RefreshToken(c context.Context, req *http.Request, refreshToken string) (*model.AuthResponse, error) {
	// Validate request token
	token, err := biz.jwtProvider.ValidateToken(refreshToken)

	if err != nil || token == nil {
		return nil, err
	}

	// Extract email from token
	claims, err := biz.jwtProvider.ParseToken(refreshToken)
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
	_, err = biz.jwtProvider.CompareToken(prevRefreshToken, refreshToken)

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

	accessToken, expAcTokenSecs, err := biz.jwtProvider.GenerateAccessToken(user.Email, model.IsAdmin(user.RoleCode))

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
