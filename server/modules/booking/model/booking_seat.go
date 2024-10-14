package model

import (
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
)

const (
	BookingSeatEntityName = "booking_seat"
)

type BookingSeat struct {
	Booking Booking    `json:"-" gorm:"column:booking;foreignKey:id"`
	Seat    model.Seat `json:"-" gorm:"column:seat;foreignKey:id"`
}

func (BookingSeat) TableName() string { return "booking_seat" }

type BookingSeatCreation struct {
	BookingId int `json:"-" gorm:"column:booking_id;foreignKey:id"`
	SeatId    int `json:"-" gorm:"column:seat_id;foreignKey:id"`
}
