package business

import (
	"context"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
)

type CreateMovieStorage interface {
	CreateMovie(ctx context.Context, data *model.MovieCreation) error
}

type createMovieBiz struct {
	storage CreateMovieStorage
}

func NewCreateMovieBiz(storage CreateMovieStorage) *createMovieBiz {
	return &createMovieBiz{storage}
}

func (biz *createMovieBiz) CreateNewMovie(ctx context.Context, data *model.MovieCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrMovieTitleIsBlank
	}

	if err := biz.storage.CreateMovie(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.MovieEntityName)
	}

	return nil
}
