package model

import (
	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	PlaceEntityName = "place"
)

type Place struct {
	common.SQLModel
	PlaceId   int    `json:"place_id" gorm:"column:place_id;primaryKey"`
	PlaceName string `json:"place_name" gorm:"column:place_name"`
	PlaceSlug string `json:"place_slug" gorm:"column:place_slug"`
}

func (Place) TableName() string { return "place" }

type PlaceCreation struct {
	PlaceId   int    `gorm:"column:place_id;primaryKey"`
	PlaceName string `json:"place_name" gorm:"column:place_name"`
	PlaceSlug string `gorm:"column:place_slug"`
}

func (PlaceCreation) TableName() string { return Place{}.TableName() }

type PlaceUpdate struct {
	PlaceName string `json:"place_name" gorm:"column:place_name"`
	PlaceSlug string `json:"place_slug" gorm:"column:place_slug"`
}

func (PlaceUpdate) TableName() string { return Place{}.TableName() }
