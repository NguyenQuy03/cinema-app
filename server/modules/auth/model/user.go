package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	UserEntityName = "user"
)

type User struct {
	common.SQLModel
	Email       string    `json:"email" gorm:"column:email"`
	Password    string    `json:"password" gorm:"column:password"`
	FullName    string    `json:"full_name" gorm:"column:full_name"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	RoleCode    string    `json:"-" gorm:"column:role_code"`
	UserRole    *UserRole `json:"role" gorm:"foreignKey:role_code;references:role_code"`
}

func (User) TableName() string { return "user" }

type UserRegister struct {
	Id          int    `gorm:"column:id;primaryKey"`
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
