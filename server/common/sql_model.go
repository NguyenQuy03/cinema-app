package common

import (
	"time"

	"gorm.io/gorm"
)

type SQLModel struct {
	Id        int            `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
