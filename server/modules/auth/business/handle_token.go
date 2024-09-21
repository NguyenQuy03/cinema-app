package business

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type HandleTokenBiz struct {
}

func (handler *HandleTokenBiz) GenerateAccessToken(user *model.User) (string, error) {
	return buildToken(user, model.AccessTokenMaxAge)
}

func (handler *HandleTokenBiz) GenerateRefreshToken(user *model.User) (string, error) {
	return buildToken(user, model.RefreshTokenMaxAge)
}

func (handler *HandleTokenBiz) ValidateToken(tokenString string) (*jwt.Token, error) {

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

func (handler *HandleTokenBiz) ExtractEmail(token *jwt.Token) (string, error) {
	// Check if the token is valid and has claims
	if token == nil {
		return "", model.ErrNilToken
	}

	// Ensure the token has claims
	claims, ok := token.Claims.(jwt.MapClaims)
	// claims, ok := token.Claims.(*model.Token)
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

func buildToken(user *model.User, expTime time.Time) (string, error) {
	claims := model.Token{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(user.UserId),
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", model.ErrGenerateToken
	}

	return tokenString, nil
}
