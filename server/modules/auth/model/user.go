package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	UserEntityName = "user"
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
)

type User struct {
	common.SQLModel
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"password" gorm:"column:password"`
	FullName    string `json:"full_name" gorm:"column:full_name"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Role        string `json:"role" gorm:"column:role"`
}

func (User) TableName() string { return "users" }

type UserRegister struct {
	Id          int    `json:"-" gorm:"column:id"`
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"password" gorm:"column:password"`
	FullName    string `json:"full_name" gorm:"column:full_name"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
}

func (UserRegister) TableName() string { return User{}.TableName() }

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string { return User{}.TableName() }
