package business

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/utils/cookieUtil"
	"github.com/NguyenQuy03/cinema-app/server/utils/jwtUtil"
)

type UpdateSessionStorage interface {
	GetUserSession(ctx context.Context, email string) (map[string]string, error)
	StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration time.Duration) error
}

type refreshTokenBiz struct {
	getUserStorage       GetUserStorage
	updateSessionStorage UpdateSessionStorage
}

func NewRefreshTokenBiz(getUserStorage GetUserStorage, updateSessionStorage UpdateSessionStorage) *refreshTokenBiz {
	return &refreshTokenBiz{
		getUserStorage:       getUserStorage,
		updateSessionStorage: updateSessionStorage,
	}
}

func (biz *refreshTokenBiz) RefreshToken(c context.Context, req *http.Request, authResponse *model.AuthResponse) error {

	// Get refresh_token from cookie
	refreshToken, err := cookieUtil.GetCookie(req, "refresh_token")

	if err != nil {
		return err
	}

	if refreshToken == "" {
		return model.ErrRequireLogin
	}

	// Validate token
	token, err := jwtUtil.ValidateToken(refreshToken)

	if err != nil || token == nil {
		return err
	}

	// Extract email from token
	email, err := jwtUtil.ExtractEmail(token)

	if err != nil {
		return err
	}

	// Get prev token in redis
	userSession, err := biz.updateSessionStorage.GetUserSession(c, email)

	if err != nil {
		return err
	}

	// Validate with prev refresh token
	if !strings.EqualFold(userSession["refresh_token"], refreshToken) {
		return model.ErrInvalidToken
	}

	user, err := biz.getUserStorage.GetUser(c, map[string]interface{}{
		"email": email,
	})

	if err != nil {
		return err
	}

	accessToken, err := jwtUtil.GenerateAccessToken(user)

	if err != nil {
		return err
	}

	// Update token in redis and return new token for user
	biz.updateSessionStorage.StoreUserSession(c, email, map[string]interface{}{"access_token": accessToken}, 0)

	authResponse.AccessToken = accessToken

	return nil
}
