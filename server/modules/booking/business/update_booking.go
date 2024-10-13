package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

type UpdateBookingStorage interface {
	UpdateBooking(ctx context.Context, conds map[string]interface{}, newData *model.BookingUpdate) error
	GetBooking(ctx context.Context, conds map[string]interface{}) (*model.Booking, error)
}

type updateBookingBiz struct {
	storage UpdateBookingStorage
}

func NewUpdateBookingBiz(storage UpdateBookingStorage) *updateBookingBiz {
	return &updateBookingBiz{storage}
}

func (biz *updateBookingBiz) UpdateBooking(ctx context.Context, id int, newData *model.BookingUpdate) error {
	oldData, err := biz.storage.GetBooking(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.BookingEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.BookingEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("booking not found"), model.BookingEntityName)
	}

	if err := biz.storage.UpdateBooking(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.BookingEntityName)
	}

	return nil
}
