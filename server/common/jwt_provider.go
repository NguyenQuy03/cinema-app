package common

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RefreshToken = "refresh_token"
	AccessToken  = "access_token"

	expireAccessTokenInSeconds  = 60 * 30      // 30 mins
	expireRefreshTokenInSeconds = 60 * 60 * 24 // 24 hours
)

var (
	ErrGenerateToken = ErrInternal(errors.New("failed to generate token"))

	ErrInvalidToken = NewUnauthorized(errors.New("token is invalid"), "The token provided is invalid", "TOKEN_INVALID_ERR")

	ErrMalformedToken = NewUnauthorized(
		errors.New("malformed token"),
		"The provided token is malformed",
		"TOKEN_MALFORMED_ERR",
	)

	ErrExpiredToken = NewUnauthorized(
		errors.New("token is expired"),
		"The provided token is expired",
		"TOKEN_EXPIRED_ERR",
	)

	ErrNilToken = NewUnauthorized(
		errors.New("token is nil"),
		"The provided token is missing or null",
		"TOKEN_NIL_ERR",
	)

	ErrClaimTypeAssertion = NewUnauthorized(
		errors.New("claims type assertion failed"),
		"The token claims type is invalid",
		"CLAIMS_TYPE_ASSERTION_ERR",
	)

	ErrRequireLogin = NewUnauthorized(
		errors.New("empty token"),
		"Session expired. Please re-login",
		"EMPTY_TOKEN",
	)
)

type JWTProvider struct{}

func (provider *JWTProvider) GenerateAccessToken(sub string) (string, int, error) {
	return provider.issueToken(sub, expireAccessTokenInSeconds)
}

func (provider *JWTProvider) GenerateRefreshToken(sub string) (string, int, error) {
	return provider.issueToken(sub, expireRefreshTokenInSeconds)
}

func (provider *JWTProvider) issueToken(sub string, expTime int) (string, int, error) {
	claims := jwt.RegisteredClaims{
		Subject:   sub,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expTime) * time.Second)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", 0, ErrGenerateToken
	}

	return tokenString, expTime, nil
}

func (provider *JWTProvider) ValidateToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check algorithm used for generating token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, NewUnauthorized(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]), "unexpected signing method", "SIGNING_METHOD_ERR")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		var returnErr error

		switch {
		case errors.Is(err, jwt.ErrSignatureInvalid):
			returnErr = ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenMalformed):
			returnErr = ErrMalformedToken
		case errors.Is(err, jwt.ErrTokenExpired):
			returnErr = ErrExpiredToken
		default:
			// Generic error handling for other cases
			returnErr = ErrInvalidToken
		}

		return nil, returnErr
	}

	return token, nil
}

func (provider *JWTProvider) ParseToken(tokenString string) (claims *jwt.RegisteredClaims, err error) {
	var rc jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(tokenString, &rc, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if !token.Valid {
		return nil, err
	}

	return &rc, nil
}

func (provider *JWTProvider) CompareToken(token1, token2 string) (bool, error) {
	// Handle Expired Token
	if token1 == "" || token2 == "" {
		return false, ErrRequireLogin
	}

	// Compare tokens
	if !strings.EqualFold(token1, token2) {
		return false, ErrInvalidToken
	}

	return true, nil
}
