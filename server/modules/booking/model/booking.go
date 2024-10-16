package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	userModel "github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	showingModel "github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

const (
	BookingEntityName = "booking"
)

type Booking struct {
	common.SQLModel
	ShowingId int                   `json:"-" gorm:"column:show_id"`
	Showing   *showingModel.Showing `json:"showing" gorm:"foreignKey:ShowingId;references:id"`
	UserId    int                   `json:"-" gorm:"column:user_id"`
	User      *userModel.User       `json:"user" gorm:"foreignKey:UserId;references:id"`
	Status    *BookingStatus        `json:"status" gorm:"column:status"`
}

func (Booking) TableName() string { return "booking" }

type BookingCreation struct {
	Id        int            `gorm:"column:id;primaryKey"`
	UserId    int            `json:"user_id" gorm:"column:user_id"`
	ShowingId int            `json:"show_id" gorm:"column:show_id"`
	Status    *BookingStatus `json:"status" gorm:"column:status;type:text"`
	Seats     []int          `json:"seats"`
	Tickets   []ticket       `json:"tickets"`
}

func (BookingCreation) TableName() string { return Booking{}.TableName() }

type BookingUpdate struct {
	Showing *showingModel.Showing `json:"showing" gorm:"column:showing;foreignKey:id"`
	Status  *BookingStatus        `json:"status" gorm:"column:status"`
}

func (BookingUpdate) TableName() string { return Booking{}.TableName() }

type ticket struct {
	TicketTypeId   int `json:"ticket_type_id"`
	TicketQuantity int `json:"ticket_quanity"`
}
