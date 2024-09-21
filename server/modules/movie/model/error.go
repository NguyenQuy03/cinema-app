package model

import (
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
)

var (
	ErrMovieDeleted      = common.NewCustomError(errors.New("movie is deleted"), "movie has been deleted", "MOVIE_DELETED_ERROR")
	ErrMovieTitleIsBlank = common.NewCustomError(errors.New("title is blank"), "title cannot be blank", "TITLE_BLANK_ERROR")
)
