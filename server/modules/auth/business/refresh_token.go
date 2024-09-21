package business

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/utils/cookieUtil"
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
	refreshToken, err := cookieUtil.GetCookie(req, model.RefreshToken)

	if err != nil {
		return err
	}

	if refreshToken == "" {
		return model.ErrRequireLogin
	}

	// Validate request token
	handleTokenBiz := HandleTokenBiz{}
	token, err := handleTokenBiz.ValidateToken(refreshToken)

	if err != nil || token == nil {
		return err
	}

	// Extract email from token
	email, err := handleTokenBiz.ExtractEmail(token)

	if err != nil {
		return err
	}

	// Get prev refresh token in redis
	userSession, err := biz.updateSessionStorage.GetUserSession(c, email)

	if err != nil {
		return err
	}

	prevRefreshToken := userSession[model.RefreshToken]

	// Handle Expired Token
	if prevRefreshToken == "" {
		return model.ErrRequireLogin
	}

	// Validate with prev refresh token
	if !strings.EqualFold(prevRefreshToken, refreshToken) {
		return model.ErrInvalidToken
	}

	// Generate and return new access token
	user, err := biz.getUserStorage.GetUser(c, map[string]interface{}{
		"email": email,
	})

	if err != nil {
		return err
	}

	accessToken, err := handleTokenBiz.GenerateAccessToken(user)

	if err != nil {
		return err
	}

	authResponse.AccessToken = accessToken

	return nil
}
