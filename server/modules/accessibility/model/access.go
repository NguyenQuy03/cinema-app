package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	AccessEntityName = "accessibility"
)

type Accessibility struct {
	common.SQLModel
	AccId      int    `json:"acc_id" gorm:"column:acc_id;primaryKey"`
	AccFeature string `json:"acc_feature" gorm:"column:acc_feature"`
	Shorten    string `json:"shorten" gorm:"column:shorten"`
}

func (Accessibility) TableName() string { return "accessibility" }

type AccessCreation struct {
	AccId      int    `gorm:"column:acc_id;primaryKey"`
	AccFeature string `json:"acc_feature" gorm:"column:acc_feature"`
	Shorten    string `json:"shorten" gorm:"column:shorten"`
}

func (AccessCreation) TableName() string { return Accessibility{}.TableName() }

type AccessUpdate struct {
	AccFeature string `json:"acc_feature" gorm:"column:acc_feature"`
	Shorten    string `json:"shorten" gorm:"column:shorten"`
}

func (AccessUpdate) TableName() string { return Accessibility{}.TableName() }
