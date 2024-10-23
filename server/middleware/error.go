package middleware

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	ErrUnauthoried = common.NewUnauthorized(
		errors.New("missing or invalid token"),
		"authorization header is missing or invalid",
		"TOKEN_MISSING_OR_INVALID_ERR",
	)

	ErrNoPermission = common.NewUnauthorized(
		errors.New("no permission"),
		"user does not have permission to access this resource",
		"NO_PERMISSION_ERR",
	)
)
