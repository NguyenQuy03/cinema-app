package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
)

type UpdateSeatStorage interface {
	UpdateSeat(ctx context.Context, conds map[string]interface{}, newData *model.SeatUpdate) error
	GetSeat(ctx context.Context, conds map[string]interface{}) (*model.Seat, error)
}

type updateSeatBiz struct {
	storage UpdateSeatStorage
}

func NewUpdateSeatBiz(storage UpdateSeatStorage) *updateSeatBiz {
	return &updateSeatBiz{storage}
}

func (biz *updateSeatBiz) UpdateSeat(ctx context.Context, id int, newData *model.SeatUpdate) error {
	oldData, err := biz.storage.GetSeat(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.SeatEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.SeatEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("seat not found"), model.SeatEntityName)
	}

	if err := biz.storage.UpdateSeat(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.SeatEntityName)
	}

	return nil
}
