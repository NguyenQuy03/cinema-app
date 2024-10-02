package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
	directorModel "github.com/NguyenQuy03/cinema-app/server/modules/director/model"
	genreModel "github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

const (
	MovieEntityName = "movie"
)

type Movie struct {
	common.SQLModel
	MovieId     int                    `json:"movie_id" gorm:"column:movie_id;primaryKey"`
	Director    directorModel.Director `json:"director" gorm:"column:director;foreignKey:DirectorId"`
	Genres      []genreModel.Genre     `json:"genres" gorm:"many2many:movie_genre;foreignKey:MovieId;joinForeignKey:MovieId;References:GenreId;joinReferences:GenreId"`
	Title       string                 `json:"title" gorm:"column:title"`
	Status      *MovieStatus           `json:"status" gorm:"column:status"`
	Duration    int                    `json:"duration" gorm:"column:duration"`
	TrailerLink string                 `json:"trailer_link" gorm:"column:trailer_link"`
	Description string                 `json:"description" gorm:"column:description"`
	// PosterImg   string                 `json:"poster_img" gorm:"column:poster_img"`
	// HeaderImg   string                 `json:"header_img" gorm:"column:header_img"`
	// AgeRating   int                    `json:"age_rating" gorm:"column:age_rating"`
	// ReleaseDate time.Time `json:"release_date" gorm:"column:release_date"`
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
