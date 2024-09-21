package model

import "github.com/golang-jwt/jwt/v5"

const (
	RefreshToken = "refresh_token"
	AccessToken  = "access_token"
)

type Token struct {
	jwt.RegisteredClaims
}
