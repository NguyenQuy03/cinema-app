package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

type ListBookingStorage interface {
	ListBooking(ctx context.Context, p *common.Paging, params ...string) ([]model.Booking, error)
}

type listBookingBiz struct {
	storage ListBookingStorage
}

func NewListBookingBiz(storage ListBookingStorage) *listBookingBiz {
	return &listBookingBiz{
		storage: storage,
	}
}

func (biz *listBookingBiz) ListBooking(ctx context.Context, p *common.Paging, params ...string) ([]model.Booking, error) {
	result, err := biz.storage.ListBooking(ctx, p)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.BookingEntityName)
	}

	return result, nil
}
