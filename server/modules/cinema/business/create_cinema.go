package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
)

type CreateCinemaStorage interface {
	CreateCinema(ctx context.Context, data *model.CinemaCreation) error
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createCinemaBiz struct {
	storage      CreateCinemaStorage
	slugProvider SlugProvider
}

func NewCreateCinemaBiz(storage CreateCinemaStorage, slugProvider SlugProvider) *createCinemaBiz {
	return &createCinemaBiz{
		storage:      storage,
		slugProvider: slugProvider,
	}
}

func (biz *createCinemaBiz) CreateCinema(ctx context.Context, data *model.CinemaCreation) error {

	data.CinemaSlug = biz.slugProvider.GenerateSlug(data.CinemaName)

	if err := biz.storage.CreateCinema(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.CinemaEntityName)
	}

	return nil
}
