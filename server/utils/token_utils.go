package utils

import (
	"os"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/modules/user/model"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role,omitempty"`
}

var ()

func buildToken(user *model.User, expTime time.Time) (string, error) {
	claims := jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
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

func GenerateAccessToken(user *model.User) (string, error) {
	token, err := buildToken(user, model.AccessTokenMaxAge)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(user *model.User) (string, error) {
	token, err := buildToken(user, model.RefreshTokenMaxAge)

	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractEmail(token *jwt.Token) (string, error) {
	// Check if the token is valid and has claims
	if token == nil {
		return "", model.ErrNilToken
	}

	// Ensure the token has claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", model.ErrClaimTypeAssertion
	}

	// Check if the token is valid
	if !token.Valid {
		return "", model.ErrInvalidToken
	}

	// Extract the email from the "sub" claim
	return claims.GetSubject()
}
