package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
)

type CreateSeatStorage interface {
	CreateSeat(ctx context.Context, data *model.SeatCreation) error
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createSeatBiz struct {
	storage CreateSeatStorage
}

func NewCreateSeatBiz(storage CreateSeatStorage) *createSeatBiz {
	return &createSeatBiz{
		storage: storage,
	}
}

func (biz *createSeatBiz) CreateSeat(ctx context.Context, data *model.SeatCreation) error {

	if err := biz.storage.CreateSeat(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.SeatEntityName)
	}

	return nil
}
