package model

import (
	"errors"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	AccessTokenMaxAge  = time.Now().Add(time.Minute * 30)
	RefreshTokenMaxAge = time.Now().Add(time.Hour * 24 * 30)

	ErrGenerateToken = common.ErrInternal(errors.New("failed to generate token"))

	ErrInvalidToken       = common.NewUnauthorized(errors.New("token is invalid"), "token is invalid", "token is invalid", "TOKEN_INVALID_ERR")
	ErrMalformedToken     = common.NewUnauthorized(errors.New("malformed token"), "Token is malformed", "The provided token is malformed", "TOKEN_MALFORMED_ERR")
	ErrExpirededToken     = common.NewUnauthorized(errors.New("token is expired"), "Token is expired", "The provided token has expired", "TOKEN_EXPIRED_ERR")
	ErrNilToken           = common.NewUnauthorized(errors.New("token is nil"), "token is nil", "token is nil", "TOKEN_NIL_ERR")
	ErrClaimTypeAssertion = common.NewUnauthorized(errors.New("claims type assertion failed"), "invalid claims type", "token claims type assertion failed", "CLAIMS_TYPE_ASSERTION_ERR")
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"-"`
}
