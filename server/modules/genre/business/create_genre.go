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

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createGenreBiz struct {
	storage      CreateGenreStorage
	slugProvider SlugProvider
}

func NewCreateGenreBiz(storage CreateGenreStorage, slugProvider SlugProvider) *createGenreBiz {
	return &createGenreBiz{
		storage:      storage,
		slugProvider: slugProvider,
	}
}

func (biz *createGenreBiz) CreateGenre(ctx context.Context, data *model.GenreCreation) error {
	name := strings.TrimSpace(data.GenreName)

	if name == "" {
		return model.ErrGenreNameIsBlank
	}

	data.GenreSlug = biz.slugProvider.GenerateSlug(data.GenreName)

	if err := biz.storage.CreateGenre(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.GenreEntityName)
	}

	return nil
}
