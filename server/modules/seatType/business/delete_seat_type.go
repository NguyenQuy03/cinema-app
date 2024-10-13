package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
)

type DeleteSeatTypeStorage interface {
	GetSeatType(ctx context.Context, conds map[string]interface{}) (*model.SeatType, error)
	DeleteSeatType(ctx context.Context, conds map[string]interface{}) error
}

type deleteSeatTypeBiz struct {
	storage DeleteSeatTypeStorage
}

func NewDeleteSeatTypeBiz(storage DeleteSeatTypeStorage) *deleteSeatTypeBiz {
	return &deleteSeatTypeBiz{
		storage: storage,
	}
}

func (biz *deleteSeatTypeBiz) DeleteSeatTypeById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetSeatType(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.SeatTypeEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.SeatTypeEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("seat type not found"), model.SeatTypeEntityName)
	}

	if err := biz.storage.DeleteSeatType(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.SeatTypeEntityName)
	}

	return nil
}
