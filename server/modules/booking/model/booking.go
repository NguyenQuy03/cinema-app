package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

const (
	BookingEntityName = "booking"
)

type Booking struct {
	common.SQLModel
	Showing model.Showing  `json:"showing" gorm:"column:showing;foreignKey:id"`
	Status  *BookingStatus `json:"status" gorm:"column:status"`
}

func (Booking) TableName() string { return "booking" }

type BookingCreation struct {
	Id      int            `gorm:"column:id;primaryKey"`
	Showing model.Showing  `json:"showing" gorm:"column:showing;foreignKey:id"`
	Status  *BookingStatus `json:"status" gorm:"column:status"`
}

func (BookingCreation) TableName() string { return Booking{}.TableName() }

type BookingUpdate struct {
	Showing model.Showing  `json:"showing" gorm:"column:showing;foreignKey:id"`
	Status  *BookingStatus `json:"status" gorm:"column:status"`
}

func (BookingUpdate) TableName() string { return Booking{}.TableName() }
