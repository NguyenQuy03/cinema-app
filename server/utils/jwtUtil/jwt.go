package jwtUtil

import "github.com/golang-jwt/jwt/v5"

type jwtClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role,omitempty"`
}
