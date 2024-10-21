package model

import (
	"time"

	"github.com/NguyenQuy03/cinema-app/server/common"
	castMemModel "github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
	directorModel "github.com/NguyenQuy03/cinema-app/server/modules/director/model"
	genreModel "github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

const (
	MovieEntityName = "movie"
)

type Movie struct {
	common.SQLModel
	DirectorID    int                       `json:"-" gorm:"column:director_id"`
	Director      *directorModel.Director   `json:"director" gorm:"foreignKey:director_id;references:Id"`
	Genres        []genreModel.Genre        `json:"genres" gorm:"many2many:movie_genre;foreignKey:Id;joinForeignKey:MovieId;References:Id;joinReferences:GenreId"`
	CastMembers   []castMemModel.CastMember `json:"cast_members" gorm:"many2many:movie_cast;foreignKey:Id;joinForeignKey:MovieId;References:Id;joinReferences:CastId"`
	Title         string                    `json:"title" gorm:"column:title"`
	Status        *MovieStatus              `json:"status" gorm:"column:status"`
	Duration      int                       `json:"duration" gorm:"column:duration"`
	TrailerLink   string                    `json:"trailer_link" gorm:"column:trailer_link"`
	Description   string                    `json:"description" gorm:"column:description"`
	PosterImg     string                    `json:"poster_img" gorm:"column:poster_img"`
	HeaderImg     string                    `json:"header_img" gorm:"column:header_img"`
	AgeRatingCode string                    `json:"age_rating" gorm:"column:age_rating"`
	ReleaseDate   time.Time                 `json:"release_date" gorm:"column:release_date"`
}

func (Movie) TableName() string { return "movie" }

type MovieCreation struct {
	Id          int                       `json:"-" gorm:"column:id;primaryKey"`
	DirectorID  int                       `json:"director_id" gorm:"column:director_id"`
	Genres      []genreModel.Genre        `json:"genres" gorm:"many2many:movie_genre;foreignKey:Id;joinForeignKey:MovieId;References:Id;joinReferences:GenreId"`
	CastMembers []castMemModel.CastMember `json:"cast_members" gorm:"many2many:movie_cast;foreignKey:Id;joinForeignKey:MovieId;References:Id;joinReferences:CastId"`
	Title       string                    `json:"title" gorm:"column:title"`
	Status      *MovieStatus              `json:"status" gorm:"column:status"`
	Duration    int                       `json:"duration" gorm:"column:duration"`
	TrailerLink string                    `json:"trailer_link" gorm:"column:trailer_link"`
	Description string                    `json:"description" gorm:"column:description"`
	PosterImg   string                    `json:"poster_img" gorm:"column:poster_img"`
	HeaderImg   string                    `json:"header_img" gorm:"column:header_img"`
	AgeRating   int                       `json:"age_rating" gorm:"column:age_rating"`
	ReleaseDate time.Time                 `json:"release_date" gorm:"column:release_date"`
}

func (MovieCreation) TableName() string { return Movie{}.TableName() }

type MovieUpdate struct {
	Id          int                       `json:"-" gorm:"column:id;primaryKey"`
	DirectorID  int                       `json:"director_id" gorm:"column:director_id"`
	Genres      []genreModel.Genre        `json:"genres" gorm:"many2many:movie_genre;joinForeignKey:Id;References:Id;joinReferences:genreId"` // Corrected joinForeignKey and joinReferences
	CastMembers []castMemModel.CastMember `json:"cast_members" gorm:"many2many:movie_cast;foreignKey:Id;joinForeignKey:MovieId;References:Id;joinReferences:CastId"`
	Title       string                    `json:"title" gorm:"column:title"`
	Status      *MovieStatus              `json:"status" gorm:"column:status"`
	Duration    int                       `json:"duration" gorm:"column:duration"`
	TrailerLink string                    `json:"trailer_link" gorm:"column:trailer_link"`
	Description string                    `json:"description" gorm:"column:description"`
	PosterImg   string                    `json:"poster_img" gorm:"column:poster_img"`
	HeaderImg   string                    `json:"header_img" gorm:"column:header_img"`
	AgeRating   int                       `json:"age_rating" gorm:"column:age_rating"`
	ReleaseDate time.Time                 `json:"release_date" gorm:"column:release_date"`
}

func (MovieUpdate) TableName() string { return Movie{}.TableName() }
