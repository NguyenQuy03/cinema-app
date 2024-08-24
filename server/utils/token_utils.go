package utils

import (
	"os"

	"github.com/NguyenQuy03/cinema-app/server/modules/user/model"
	"github.com/golang-jwt/jwt/v5"
)

func buildToken(user *model.User, expTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"exp": expTime,
	})

	stringValue, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", model.ErrGenerateToken
	}

	return stringValue, nil
}

func GenerateAccessToken(user *model.User) (string, error) {
	token, err := buildToken(user, model.AccessTokenMaxAge)

	if err != nil {
		return "", nil
	}

	return token, nil
}

func GenerateRefreshToken(user *model.User) (string, error) {
	token, err := buildToken(user, model.RefreshTokenMaxAge)

	if err != nil {
		return "", nil
	}

	return token, nil
}
