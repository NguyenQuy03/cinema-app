package jwtUtil

import (
	"os"
	"strconv"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

func buildToken(user *model.User, expTime time.Time) (string, error) {
	claims := jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(user.Id),
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
		Role: user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", model.ErrGenerateToken
	}

	return tokenString, nil
}
