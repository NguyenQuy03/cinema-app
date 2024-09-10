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

	ErrInvalidToken = common.NewUnauthorized(
		errors.New("token is invalid"),
		"The token provided is invalid",
		"TOKEN_INVALID_ERR",
	)

	ErrMalformedToken = common.NewUnauthorized(
		errors.New("malformed token"),
		"The provided token is malformed",
		"TOKEN_MALFORMED_ERR",
	)

	ErrExpiredToken = common.NewUnauthorized(
		errors.New("token is expired"),
		"The provided token is expired",
		"TOKEN_EXPIRED_ERR",
	)

	ErrNilToken = common.NewUnauthorized(
		errors.New("token is nil"),
		"The provided token is missing or null",
		"TOKEN_NIL_ERR",
	)

	ErrClaimTypeAssertion = common.NewUnauthorized(
		errors.New("claims type assertion failed"),
		"The token claims type is invalid",
		"CLAIMS_TYPE_ASSERTION_ERR",
	)

	ErrRequireLogin = common.NewUnauthorized(
		errors.New("empty token"),
		"Session expired. Please re-login",
		"EMPTY_TOKEN",
	)
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"-"`
}
