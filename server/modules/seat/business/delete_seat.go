package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
)

type DeleteSeatStorage interface {
	GetSeat(ctx context.Context, conds map[string]interface{}) (*model.Seat, error)
	DeleteSeat(ctx context.Context, conds map[string]interface{}) error
}

type deleteSeatBiz struct {
	storage DeleteSeatStorage
}

func NewDeleteSeatBiz(storage DeleteSeatStorage) *deleteSeatBiz {
	return &deleteSeatBiz{
		storage: storage,
	}
}

func (biz *deleteSeatBiz) DeleteSeatById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetSeat(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.SeatEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.SeatEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("seat not found"), model.SeatEntityName)
	}

	if err := biz.storage.DeleteSeat(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.SeatEntityName)
	}

	return nil
}
