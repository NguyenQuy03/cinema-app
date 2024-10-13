package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

type CreateBookingStorage interface {
	CreateBooking(ctx context.Context, data *model.BookingCreation) error
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createBookingBiz struct {
	storage CreateBookingStorage
}

func NewCreateBookingBiz(storage CreateBookingStorage) *createBookingBiz {
	return &createBookingBiz{
		storage: storage,
	}
}

func (biz *createBookingBiz) CreateBooking(ctx context.Context, data *model.BookingCreation) error {
	if err := biz.storage.CreateBooking(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.BookingEntityName)
	}

	return nil
}
