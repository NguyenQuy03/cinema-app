package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

const (
	DirectorEntityName = "director"
)

var (
	ErrDirectorNameIsBlank = common.NewCustomError(errors.New("director name is blank"), "director name cannot be blank", "NAME_BLANK_ERROR")
	ErrDirectorNotFound    = common.ErrEntityNotFound(errors.New("director not found"), DirectorEntityName)
)

type Director struct {
	DirectorId   int    `json:"director_id" gorm:"director_id;primaryKey"`
	DirectorName string `json:"director_name" gorm:"director_name"`
}

func (Director) TableName() string { return "director" }

type DirectorCreation struct {
	DirectorId   int    `gorm:"director_id;primaryKey"`
	DirectorName string `json:"director_name" gorm:"director_name"`
}

func (DirectorCreation) TableName() string { return Director{}.TableName() }

type DirectorUpdate struct {
	DirectorName string `json:"director_name" gorm:"director_name"`
}

func (DirectorUpdate) TableName() string { return Director{}.TableName() }
