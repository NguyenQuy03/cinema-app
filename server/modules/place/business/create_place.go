package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
)

type CreatePlaceStorage interface {
	CreatePlace(ctx context.Context, data *model.PlaceCreation) error
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createPlaceBiz struct {
	storage      CreatePlaceStorage
	slugProvider SlugProvider
}

func NewCreatePlaceBiz(storage CreatePlaceStorage, slugProvider SlugProvider) *createPlaceBiz {
	return &createPlaceBiz{
		storage:      storage,
		slugProvider: slugProvider,
	}
}

func (biz *createPlaceBiz) CreatePlace(ctx context.Context, data *model.PlaceCreation) error {

	data.PlaceSlug = biz.slugProvider.GenerateSlug(data.PlaceName)

	if err := biz.storage.CreatePlace(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.PlaceEntityName)
	}

	return nil
}
