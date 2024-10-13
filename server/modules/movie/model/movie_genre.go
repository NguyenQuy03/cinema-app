package model

type MovieGenre struct {
	MovieId int `gorm:"column:movie_id"`
	GenreId int `gorm:"column:genre_id"`
}

func (MovieGenre) TableName() string {
	return "movie_genre"
}
