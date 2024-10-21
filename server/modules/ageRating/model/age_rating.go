package model

const (
	AgeRatingEntityName = "age_rating"
)

type AgeRating struct {
	RatingCode  string `json:"rating_code" gorm:"column:rating_code;primaryKey"`
	RatingValue int    `json:"rating_value" gorm:"co:age_rating_value"`
	Description string `json:"description" gorm:"column:description"`
}

func (AgeRating) TableName() string { return "age_rating" }

type AgeRatingCreation struct {
	RatingCode  string `json:"rating_code" gorm:"column:rating_code;primaryKey"`
	RatingValue int    `json:"rating_value" gorm:"co:age_rating_value"`
	Description string `json:"description" gorm:"column:description"`
}

func (AgeRatingCreation) TableName() string { return AgeRating{}.TableName() }

type AgeRatingUpdate struct {
	RatingCode  string `json:"rating_code" gorm:"column:rating_code"`
	RatingValue int    `json:"rating_value" gorm:"co:age_rating_value"`
	Description string `json:"description" gorm:"column:description"`
}

func (AgeRatingUpdate) TableName() string { return AgeRating{}.TableName() }
