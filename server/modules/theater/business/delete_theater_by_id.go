package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

type DeleteTheaterStorage interface {
	GetTheater(ctx context.Context, conds map[string]interface{}) (*model.Theater, error)
	DeleteTheater(ctx context.Context, conds map[string]interface{}) error
}

type deleteTheaterBiz struct {
	storage DeleteTheaterStorage
}

func NewDeleteTheaterBiz(storage DeleteTheaterStorage) *deleteTheaterBiz {
	return &deleteTheaterBiz{storage}
}

func (biz *deleteTheaterBiz) DeleteTheaterById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetTheater(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.TheaterEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.TheaterEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("theater not found"), model.TheaterEntityName)
	}

	if err := biz.storage.DeleteTheater(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.TheaterEntityName)
	}

	return nil
}
