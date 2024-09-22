package business

import (
	"context"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

type CreateGenreStorage interface {
	CreateGenre(ctx context.Context, data *model.GenreCreation) error
}

type createGenreBiz struct {
	storage CreateGenreStorage
}

func NewCreateGenreBiz(storage CreateGenreStorage) *createGenreBiz {
	return &createGenreBiz{storage}
}

func (biz *createGenreBiz) CreateGenre(ctx context.Context, data *model.GenreCreation) error {
	title := strings.TrimSpace(data.GenreName)

	if title == "" {
		return model.ErrGenreNameIsBlank
	}

	if err := biz.storage.CreateGenre(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.GenreEntityName)
	}

	return nil
}
