package model

import (
	"time"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	CastMemberEntityName = "cast_member"
)

type CastMember struct {
	common.SQLModel
	CastName string    `json:"cast_name" gorm:"cast_name"`
	Dob      time.Time `json:"dob" gorm:"dob"`
	Bio      string    `json:"bio" gorm:"bio"`
}

func (CastMember) TableName() string { return "cast_member" }

type CastMemberCreation struct {
	Id       int       `gorm:"id;primaryKey"`
	CastName string    `json:"cast_name" gorm:"cast_name"`
	Dob      time.Time `json:"dob" gorm:"dob"`
	Bio      string    `json:"bio" gorm:"bio"`
}

func (CastMemberCreation) TableName() string { return CastMember{}.TableName() }

type CastMemberUpdate struct {
	CastName string    `json:"cast_name" gorm:"cast_name"`
	Dob      time.Time `json:"dob" gorm:"dob"`
	Bio      string    `json:"bio" gorm:"bio"`
}

func (CastMemberUpdate) TableName() string { return CastMember{}.TableName() }
