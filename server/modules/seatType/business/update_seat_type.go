package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
)

type UpdateSeatTypeStorage interface {
	UpdateSeatType(ctx context.Context, conds map[string]interface{}, newData *model.SeatTypeUpdate) error
	GetSeatType(ctx context.Context, conds map[string]interface{}) (*model.SeatType, error)
}

type updateSeatTypeBiz struct {
	storage UpdateSeatTypeStorage
}

func NewUpdateSeatTypeBiz(storage UpdateSeatTypeStorage) *updateSeatTypeBiz {
	return &updateSeatTypeBiz{storage}
}

func (biz *updateSeatTypeBiz) UpdateSeatType(ctx context.Context, id int, newData *model.SeatTypeUpdate) error {
	oldData, err := biz.storage.GetSeatType(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.SeatTypeEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.SeatTypeEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("seat type not found"), model.SeatTypeEntityName)
	}

	if err := biz.storage.UpdateSeatType(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.SeatTypeEntityName)
	}

	return nil
}
