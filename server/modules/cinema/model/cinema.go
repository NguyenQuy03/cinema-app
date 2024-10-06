package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
)

const (
	CinemaEntityName = "cinema"
)

type Cinema struct {
	common.SQLModel
	Place      model.Place `json:"place" gorm:"column:place;foreignKey:PlaceId"`
	CinemaName string      `json:"cinema_name" gorm:"column:cinema_name"`
	CinemaSlug string      `json:"cinema_slug" gorm:"column:cinema_slug"`
}

func (Cinema) TableName() string { return "cinema" }

type CinemaCreation struct {
	Id         int         `gorm:"column:id;primaryKey"`
	Place      model.Place `json:"place" gorm:"column:place;foreignKey:PlaceId"`
	CinemaName string      `json:"cinema_name" gorm:"column:cinema_name"`
	CinemaSlug string      `json:"-" gorm:"column:cinema_slug"`
}

func (CinemaCreation) TableName() string { return Cinema{}.TableName() }

type CinemaUpdate struct {
	Place      model.Place `json:"place_slug" gorm:"column:place;foreignKey:PlaceId"`
	CinemaName string      `json:"cinema_name" gorm:"column:cinema_name"`
	CinemaSlug string      `json:"-" gorm:"column:cinema_slug"`
}

func (CinemaUpdate) TableName() string { return Cinema{}.TableName() }
