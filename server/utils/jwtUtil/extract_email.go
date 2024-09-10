package jwtUtil

import (
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

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
