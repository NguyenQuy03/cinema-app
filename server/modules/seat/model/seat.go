package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	seatTypeModel "github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
	theaterModel "github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

const (
	SeatEntityName = "seat"
)

type Seat struct {
	common.SQLModel
	SeatType     seatTypeModel.SeatType `json:"seat_type" gorm:"column:seatType;foreignKey:id"`
	Theater      theaterModel.Theater   `json:"theater" gorm:"column:Theater;foreignKey:id"`
	SeatLocation string                 `json:"seat_location" gorm:"seat_location"`
	Status       *SeatStatus            `json:"status" gorm:"status"`
}

func (Seat) TableName() string { return "seat" }

type SeatCreation struct {
	Id           int         `gorm:"id;primaryKey"`
	SeatTypeId   int         `json:"seat_type_id" gorm:"column:seat_type_id"`
	TheaterId    int         `json:"theater_id" gorm:"column:theater_id"`
	SeatLocation string      `json:"seat_location" gorm:"seat_location"`
	Status       *SeatStatus `json:"status" gorm:"status"`
}

func (SeatCreation) TableName() string { return Seat{}.TableName() }

type SeatUpdate struct {
	SeatType     seatTypeModel.SeatType `json:"seat_type" gorm:"column:seatType;foreignKey:id"`
	Theater      theaterModel.Theater   `json:"theater" gorm:"column:Theater;foreignKey:id"`
	SeatLocation string                 `json:"seat_location" gorm:"seat_location"`
	Status       *SeatStatus            `json:"status" gorm:"status"`
}

func (SeatUpdate) TableName() string { return Seat{}.TableName() }
