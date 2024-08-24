package model

import (
	"errors"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	AccessTokenMaxAge  = time.Now().Add(time.Minute * 30).Unix()
	RefreshTokenMaxAge = time.Now().Add(time.Hour * 24 * 30).Unix()

	ErrGenerateToken = common.ErrInternal(errors.New("failed to generate token"))
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
