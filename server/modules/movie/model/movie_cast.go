package model

type MovieCast struct {
	MovieId int `gorm:"column:movie_id"`
	CastId  int `gorm:"column:cast_id"`
}

func (MovieCast) TableName() string {
	return "movie_cast"
}
