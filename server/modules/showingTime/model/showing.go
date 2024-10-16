package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	movieModel "github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
	theaterModel "github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

const (
	ShowingEntityName = "showing_time"
)

type Showing struct {
	common.SQLModel
	MovieId     int                   `json:"-" gorm:"column:movie_id"`
	Movie       *movieModel.Movie     `json:"movie" gorm:"foreignKey:movie_id;references:id"`
	TheaterId   int                   `json:"-" gorm:"column:theater_id"`
	Theater     *theaterModel.Theater `json:"theater" gorm:"foreignKey:theater_id;references:id"`
	ShowingDate string                `json:"showing_date" gorm:"column:showing_date"`
	BasePrice   int                   `json:"base_price" gorm:"column:base_price"`
}

func (Showing) TableName() string { return "showing_time" }

type ShowingCreation struct {
	Id          int                   `json:"-" gorm:"column:id;primaryKey"`
	Movie       *movieModel.Movie     `json:"movie" gorm:"column:movie;foreignKey:id"`
	Theater     *theaterModel.Theater `json:"theater" gorm:"column:theater;foreignKey:id"`
	ShowingDate string                `json:"showing_date" gorm:"column:showing_date"`
	BasePrice   int                   `json:"base_price" gorm:"column:base_price"`
}

func (ShowingCreation) TableName() string { return Showing{}.TableName() }

type ShowingUpdate struct {
	Movie       *movieModel.Movie     `json:"movie" gorm:"column:movie;foreignKey:id"`
	Theater     *theaterModel.Theater `json:"theater" gorm:"column:theater;foreignKey:id"`
	ShowingDate string                `json:"showing_date" gorm:"column:showing_date"`
	BasePrice   int                   `json:"base_price" gorm:"column:base_price"`
}

func (ShowingUpdate) TableName() string { return Showing{}.TableName() }
