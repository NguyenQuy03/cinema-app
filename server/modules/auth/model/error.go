package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	ErrEmailInvalid = common.NewCustomError(errors.New("email is invalid"), "email is invalid", "EMAIL_INVALID_ERROR")
	ErrHashPassword = common.ErrInternal(errors.New("error hash password"))
	ErrUserNotExist = common.NewCustomError(errors.New("user is not exist"), "user is not exist", "USER_NOT_EXIST_ERROR")

	ErrUserExisted        = common.NewConflict(errors.New("user is existed"), "the email has existed, please choose another", "USER_EXISTED_ERROR")
	ErrEmailOrPassMissing = common.NewCustomError(
		errors.New("missing username or password"),
		"missing username or password",
		"MISSING_UNAME_OR_PASS_ERR",
	)

	ErrShortPass    = common.NewCustomError(errors.New("short password"), "your password is too short", "SHORT_PASS_ERROR")
	ErrLoginFailure = common.NewCustomError(
		errors.New("email or password is incorrect"),
		"your email or password is incorrect",
		"LOGIN_FAILURE",
	)

	ErrRequireLogin = common.NewUnauthorized(
		errors.New("empty token"),
		"Session expired. Please re-login",
		"EMPTY_TOKEN",
	)

	ErrInvalidToken = common.NewUnauthorized(errors.New("token is invalid"), "The token provided is invalid", "TOKEN_INVALID_ERR")
)
