package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	directorModel "github.com/NguyenQuy03/cinema-app/server/modules/director/model" // Ensure this import is necessary
	genreModel "github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

const (
	MovieEntityName = "movie"
)

type Movie struct {
	common.SQLModel
	MovieId     int                    `json:"movie_id" gorm:"column:movie_id;primaryKey"`
	Title       string                 `json:"title" gorm:"column:title"`
	Description string                 `json:"description" gorm:"column:description"`
	Duration    int                    `json:"duration" gorm:"column:duration"`
	Genre       string                 `json:"genre" gorm:"column:genre"`
	TrailerLink string                 `json:"trailer_link" gorm:"column:trailer_link"`
	Status      *MovieStatus           `json:"status" gorm:"column:status"`
	Director    directorModel.Director `json:"director" gorm:"column:director;foreignKey:DirectorId"`
	Genres      []genreModel.Genre     `json:"genres" gorm:"many2many:movie_genre;foreignKey:MovieId;joinForeignKey:MovieId;References:GenreId;joinReferences:GenreId"`
}

func (Movie) TableName() string { return "movie" }

type MovieCreation struct {
	MovieId     int          `json:"-" gorm:"column:movie_id"`
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

type MovieGenre struct {
	MovieId int `json:"movie_id" gorm:"column:movie_id;primaryKey"`
	GenreId int `json:"genre_id" gorm:"column:genre_id;primaryKey"`
}

func (MovieGenre) TableName() string { return "movie_genre" }
