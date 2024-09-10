package jwtUtil

import "github.com/NguyenQuy03/cinema-app/server/modules/auth/model"

func GenerateAccessToken(user *model.User) (string, error) {
	return buildToken(user, model.AccessTokenMaxAge)
}
