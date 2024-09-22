package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	GenreEntityName = "genre"
)

var (
	// ErrMovieDeleted     = common.NewCustomError(errors.New("movie is deleted"), "movie has been deleted", "MOVIE_DELETED_ERROR")
	ErrGenreNameIsBlank = common.NewCustomError(errors.New("genre name is blank"), "genre name cannot be blank", "NAME_BLANK_ERROR")
)

type Genre struct {
	GenreId   int    `json:"genre_id" gorm:"genre_id"`
	GenreName string `json:"genre_name" gorm:"genre_name"`
}

func (Genre) TableName() string { return "genre" }

type GenreCreation struct {
	GenreId   int    `gorm:"genre_id;primaryKey"`
	GenreName string `json:"genre_name" gorm:"genre_name"`
}

func (GenreCreation) TableName() string { return Genre{}.TableName() }

type GenreUpdate struct {
	GenreName string `json:"genre_name" gorm:"genre_name"`
}

func (GenreUpdate) TableName() string { return Genre{}.TableName() }
