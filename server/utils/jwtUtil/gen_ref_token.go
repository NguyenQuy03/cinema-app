package jwtUtil

import "github.com/NguyenQuy03/cinema-app/server/modules/auth/model"

func GenerateRefreshToken(user *model.User) (string, error) {
	return buildToken(user, model.RefreshTokenMaxAge)
}
