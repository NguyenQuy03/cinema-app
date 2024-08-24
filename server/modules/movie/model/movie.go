package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	MovieEntityName = "movie"
)

var (
	ErrMovieDeleted      = common.NewCustomError(errors.New("movie is deleted"), "movie has been deleted", "MOVIE_DELETED_ERROR")
	ErrMovieTitleIsBlank = common.NewCustomError(errors.New("title is blank"), "title cannot be blank", "TITLE_BLANK_ERROR")
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
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Duration    int    `json:"duration" gorm:"column:duration"`
	Genre       string `json:"genre" gorm:"column:genre"`
	TrailerLink string `json:"trailer_link" gorm:"column:trailer_link"`
}

func (MovieUpdate) TableName() string { return Movie{}.TableName() }
