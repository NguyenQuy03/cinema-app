package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

type GetBookingStorage interface {
	GetBooking(ctx context.Context, conds map[string]interface{}) (*model.Booking, error)
}

type getBookingBiz struct {
	storage GetBookingStorage
}

func NewGetBookingBiz(storage GetBookingStorage) *getBookingBiz {
	return &getBookingBiz{storage}
}

func (biz *getBookingBiz) GetBookingById(ctx context.Context, id int) (*model.Booking, error) {
	result, err := biz.storage.GetBooking(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.BookingEntityName)
	}

	return result, nil
}
