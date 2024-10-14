package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	SeatTypeEntityName = "seat_type"
)

type SeatType struct {
	common.SQLModel
	TypeName      string `json:"type_name" gorm:"type_name"`
	Slug          string `json:"slug" gorm:"slug"`
	SeatSurcharge int    `json:"seat_surcharge" gorm:"seat_surcharge"`
	SeatRange     int    `json:"seat_range" gorm:"seat_range"`
}

func (SeatType) TableName() string { return "seat_type" }

type SeatTypeCreation struct {
	Id            int    `gorm:"id;primaryKey"`
	TypeName      string `json:"type_name" gorm:"type_name"`
	Slug          string `json:"-" gorm:"slug"`
	SeatSurcharge int    `json:"seat_surcharge" gorm:"seat_surcharge"`
	SeatRange     int    `json:"seat_range" gorm:"seat_range"`
}

func (SeatTypeCreation) TableName() string { return SeatType{}.TableName() }

type SeatTypeUpdate struct {
	TypeName      string `json:"type_name" gorm:"type_name"`
	Slug          string `json:"slug" gorm:"slug"`
	SeatSurcharge int    `json:"seat_surcharge" gorm:"seat_surcharge"`
	SeatRange     int    `json:"seat_range" gorm:"seat_range"`
}

func (SeatTypeUpdate) TableName() string { return SeatType{}.TableName() }
