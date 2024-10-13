package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

type DeleteBookingStorage interface {
	GetBooking(ctx context.Context, conds map[string]interface{}) (*model.Booking, error)
	DeleteBooking(ctx context.Context, conds map[string]interface{}) error
}

type deleteBookingBiz struct {
	storage DeleteBookingStorage
}

func NewDeleteBookingBiz(storage DeleteBookingStorage) *deleteBookingBiz {
	return &deleteBookingBiz{
		storage: storage,
	}
}

func (biz *deleteBookingBiz) DeleteBookingById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetBooking(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.BookingEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.BookingEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("booking not found"), model.BookingEntityName)
	}

	if err := biz.storage.DeleteBooking(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.BookingEntityName)
	}

	return nil
}
