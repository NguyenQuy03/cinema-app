package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	ErrInvalidInput = common.NewCustomError(errors.New("input is invalid"), "input is invalid", "INVALID_INPUT_ERR")
)
