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
	CinemaId      int                        `json:"-" gorm:"column:cinema_id"`
	Cinema        *cinemaModel.Cinema        `json:"cinema" gorm:"foreignKey:cinema_id;references:id"`
	AccessId      int                        `json:"-" gorm:"column:acc_id"`
	Accessibility *accessModel.Accessibility `json:"accessibility" gorm:"foreignKey:acc_id;references:id"`
	ExperId       int                        `json:"-" gorm:"column:exp_id"`
	Experience    *experModel.Experience     `json:"experience" gorm:"foreignKey:exp_id;references:id"`
	TheaterNum    int                        `json:"theater_num" gorm:"column:theater_num"`
}

func (Theater) TableName() string { return "theater" }

type TheaterCreation struct {
	Id            int                        `json:"-" gorm:"column:id;primaryKey"`
	Cinema        *cinemaModel.Cinema        `json:"cinema" gorm:"column:cinema;foreignKey:Id"`
	Accessibility *accessModel.Accessibility `json:"accessibility" gorm:"column:accessibility;foreignKey:Id"`
	Experience    *experModel.Experience     `json:"experience" gorm:"column:experience;foreignKey:Id"`
	TheaterNum    int                        `json:"theater_num" gorm:"column:theater_num"`
}

func (TheaterCreation) TableName() string { return Theater{}.TableName() }

type TheaterUpdate struct {
	Cinema        *cinemaModel.Cinema        `json:"cinema" gorm:"column:cinema;foreignKey:Id"`
	Accessibility *accessModel.Accessibility `json:"accessibility" gorm:"column:accessibility;foreignKey:Id"`
	Experience    *experModel.Experience     `json:"experience" gorm:"column:experience;foreignKey:Id"`
	TheaterNum    int                        `json:"theater_num" gorm:"column:theater_num"`
}

func (TheaterUpdate) TableName() string { return Theater{}.TableName() }
