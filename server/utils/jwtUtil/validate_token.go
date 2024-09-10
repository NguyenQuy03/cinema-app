package jwtUtil

import (
	"errors"
	"fmt"
	"os"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check algorithm used for generating token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, common.NewUnauthorized(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]), "unexpected signing method", "SIGNING_METHOD_ERR")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		var returnErr error

		switch {
		case errors.Is(err, jwt.ErrSignatureInvalid):
			returnErr = model.ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenMalformed):
			returnErr = model.ErrMalformedToken
		case errors.Is(err, jwt.ErrTokenExpired):
			returnErr = model.ErrExpiredToken
		default:
			// Generic error handling for other cases
			returnErr = model.ErrInvalidToken
		}

		return nil, returnErr
	}

	return token, nil
}
