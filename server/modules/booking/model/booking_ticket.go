package model

import (
	ticketTypeModel "github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

const (
	BookingTicketEntityName = "booking_ticket"
)

type BookingTicket struct {
	BookingId      Booking                    `json:"booking_id" gorm:"column:booking;foreignKey:id"`
	TicketTypeId   ticketTypeModel.TicketType `json:"ticket_type_id" gorm:"column:ticket_type_id;foreignKey:id"`
	TicketQuantity int                        `json:"ticket_quantity" gorm:"column:ticket_quantity"`
}

func (BookingTicket) TableName() string { return "booking_ticket" }
