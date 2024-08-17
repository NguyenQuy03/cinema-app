package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrMovieDeleted = errors.New("movie is deleted")
)

type Movie struct {
	common.SQLModel
	Title       string       `json:"title" gorm:"column:title"`
	Description string       `json:"description" gorm:"column:description"`
	Duration    int          `json:"duration" gorm:"column:duration"`
	Genre       string       `json:"genre" gorm:"column:genre"`
	TrailerLink string       `json:"trailer_link" gorm:"column:trailer_link"`
	Status      *MovieStatus `json:"status" gorm:"column:status"`
}

func (Movie) TableName() string { return "movies" }

type MovieCreation struct {
	Id          int          `json:"-" gorm:"column:id"`
	Title       string       `json:"title" gorm:"column:title"`
	Description string       `json:"description" gorm:"column:description"`
	Duration    int          `json:"duration" gorm:"column:duration"`
	Genre       string       `json:"genre" gorm:"column:genre"`
	TrailerLink string       `json:"trailer_link" gorm:"column:trailer_link"`
	Status      *MovieStatus `json:"status" gorm:"column:status"`
}

func (MovieCreation) TableName() string { return Movie{}.TableName() }

type MovieUpdate struct {
	Title       string  `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Duration    int     `json:"duration" gorm:"column:duration"`
	Genre       string  `json:"genre" gorm:"column:genre"`
	TrailerLink string  `json:"trailer_link" gorm:"column:trailer_link"`
}

func (MovieUpdate) TableName() string { return Movie{}.TableName() }
