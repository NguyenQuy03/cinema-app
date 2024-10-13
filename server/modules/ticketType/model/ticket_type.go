package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	TicketTypeEntityName = "ticket_type"
)

type TicketType struct {
	common.SQLModel
	TicketName      string `json:"ticket_name" gorm:"ticket_name"`
	Slug            string `json:"slug" gorm:"slug"`
	TicketSurcharge int    `json:"ticket_surcharge" gorm:"ticket_surcharge"`
}

func (TicketType) TableName() string { return "ticket_type" }

type TicketTypeCreation struct {
	Id              int    `gorm:"id;primaryKey"`
	TicketName      string `json:"ticket_name" gorm:"ticket_name"`
	Slug            string `json:"slug" gorm:"slug"`
	TicketSurcharge int    `json:"ticket_surcharge" gorm:"ticket_surcharge"`
}

func (TicketTypeCreation) TableName() string { return TicketType{}.TableName() }

type TicketTypeUpdate struct {
	TicketName      string `json:"ticket_name" gorm:"ticket_name"`
	Slug            string `json:"slug" gorm:"slug"`
	TicketSurcharge int    `json:"ticket_surcharge" gorm:"ticket_surcharge"`
}

func (TicketTypeUpdate) TableName() string { return TicketType{}.TableName() }
