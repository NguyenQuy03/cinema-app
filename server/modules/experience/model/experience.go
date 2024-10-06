package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	ExperienceEntityName = "experience"
)

type Experience struct {
	common.SQLModel
	ExpFeature  string `json:"exp_feature" gorm:"column:exp_feature"`
	Description string `json:"description" gorm:"column:description"`
}

func (Experience) TableName() string { return "experience" }

type ExperienceCreation struct {
	Id          int    `gorm:"column:id;primaryKey"`
	ExpFeature  string `json:"exp_feature" gorm:"column:exp_feature"`
	Description string `json:"description" gorm:"column:description"`
}

func (ExperienceCreation) TableName() string { return Experience{}.TableName() }

type ExperienceUpdate struct {
	ExpFeature  string `json:"exp_feature" gorm:"column:exp_feature"`
	Description string `json:"description" gorm:"column:description"`
}

func (ExperienceUpdate) TableName() string { return Experience{}.TableName() }
