package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	accessModel "github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
	cinemaModel "github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
	experModel "github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
)

const (
	TheaterEntityName = "theater"
)

type Theater struct {
	common.SQLModel
	Cinema        cinemaModel.Cinema        `json:"cinema" gorm:"column:cinema;foreignKey:CinemaId"`
	Accessibility accessModel.Accessibility `json:"accessibility" gorm:"column:accessibility;foreignKey:AccId"`
	Experience    experModel.Experience     `json:"experience" gorm:"column:experience;foreignKey:ExpId"`
	TheaterNum    string                    `json:"theater_num" gorm:"column:theater_num"`
}

func (Theater) TableName() string { return "theater" }

type TheaterCreation struct {
	Id            int                       `json:"-" gorm:"column:id;primaryKey"`
	Cinema        cinemaModel.Cinema        `json:"cinema" gorm:"column:cinema;foreignKey:CinemaId"`
	Accessibility accessModel.Accessibility `json:"accessibility" gorm:"column:accessibility;foreignKey:AccId"`
	Experience    experModel.Experience     `json:"experience" gorm:"column:experience;foreignKey:ExpId"`
	TheaterNum    string                    `json:"theater_num" gorm:"column:theater_num"`
}

func (TheaterCreation) TableName() string { return Theater{}.TableName() }

type TheaterUpdate struct {
	Cinema        cinemaModel.Cinema        `json:"cinema" gorm:"column:cinema;foreignKey:CinemaId"`
	Accessibility accessModel.Accessibility `json:"accessibility" gorm:"column:accessibility;foreignKey:AccId"`
	Experience    experModel.Experience     `json:"experience" gorm:"column:experience;foreignKey:ExpId"`
	TheaterNum    string                    `json:"theater_num" gorm:"column:theater_num"`
}

func (TheaterUpdate) TableName() string { return Theater{}.TableName() }
