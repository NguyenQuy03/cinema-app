package model

import (
	ticketTypeModel "github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

const (
	BookingTicketEntityName = "booking_ticket"
)

type BookingTicket struct {
<<<<<<< HEAD
	BookingId      Booking                    `json:"-" gorm:"column:booking;foreignKey:id"`
=======
	Booking        Booking                    `json:"-" gorm:"column:booking_id;foreignKey:id"`
>>>>>>> development
	TicketType     ticketTypeModel.TicketType `json:"-" gorm:"column:ticket_type_id;foreignKey:id"`
	TicketQuantity int                        `json:"-" gorm:"column:ticket_quanity"`
}

func (BookingTicket) TableName() string { return "booking_ticket" }

type BookingTicketCreation struct {
	BookingId      int `json:"-" gorm:"column:booking_id"`
	TicketTypeId   int `json:"-" gorm:"column:ticket_type_id"`
	TicketQuantity int `json:"-" gorm:"column:ticket_quanity"`
}
<<<<<<< HEAD
=======

func (BookingTicketCreation) TableName() string { return BookingTicket{}.TableName() }
>>>>>>> development
