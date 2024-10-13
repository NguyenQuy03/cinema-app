package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
)

type GetSeatTypeStorage interface {
	GetSeatType(ctx context.Context, conds map[string]interface{}) (*model.SeatType, error)
}

type getSeatTypeBiz struct {
	storage GetSeatTypeStorage
}

func NewGetSeatTypeBiz(storage GetSeatTypeStorage) *getSeatTypeBiz {
	return &getSeatTypeBiz{storage}
}

func (biz *getSeatTypeBiz) GetSeatTypeById(ctx context.Context, id int) (*model.SeatType, error) {
	result, err := biz.storage.GetSeatType(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.SeatTypeEntityName)
	}

	return result, nil
}
