package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
)

type GetSeatStorage interface {
	GetSeat(ctx context.Context, conds map[string]interface{}) (*model.Seat, error)
}

type getSeatBiz struct {
	storage GetSeatStorage
}

func NewGetSeatBiz(storage GetSeatStorage) *getSeatBiz {
	return &getSeatBiz{storage}
}

func (biz *getSeatBiz) GetSeatById(ctx context.Context, id int) (*model.Seat, error) {
	result, err := biz.storage.GetSeat(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.SeatEntityName)
	}

	return result, nil
}
