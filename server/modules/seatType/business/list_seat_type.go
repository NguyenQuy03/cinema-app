package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
)

type ListSeatTypeStorage interface {
	ListSeatType(ctx context.Context, paging *common.Paging, params ...string) ([]model.SeatType, error)
}

type listSeatTypeBiz struct {
	storage ListSeatTypeStorage
}

func NewListSeatTypeBiz(storage ListSeatTypeStorage) *listSeatTypeBiz {
	return &listSeatTypeBiz{storage}
}

func (biz *listSeatTypeBiz) ListSeatType(ctx context.Context, paging *common.Paging) ([]model.SeatType, error) {
	result, err := biz.storage.ListSeatType(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.SeatTypeEntityName)
	}

	return result, nil
}
