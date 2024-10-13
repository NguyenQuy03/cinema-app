package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
)

type CreateSeatTypeStorage interface {
	CreateSeatType(ctx context.Context, data *model.SeatTypeCreation) error
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createSeatTypeBiz struct {
	storage      CreateSeatTypeStorage
	slugProvider SlugProvider
}

func NewCreateSeatTypeBiz(storage CreateSeatTypeStorage, slugProvider SlugProvider) *createSeatTypeBiz {
	return &createSeatTypeBiz{
		storage:      storage,
		slugProvider: slugProvider,
	}
}

func (biz *createSeatTypeBiz) CreateSeatType(ctx context.Context, data *model.SeatTypeCreation) error {

	data.Slug = biz.slugProvider.GenerateSlug(data.TypeName)

	if err := biz.storage.CreateSeatType(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.SeatTypeEntityName)
	}

	return nil
}
