package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

type DeleteShowingStorage interface {
	GetShowing(ctx context.Context, conds map[string]interface{}) (*model.Showing, error)
	DeleteShowing(ctx context.Context, conds map[string]interface{}) error
}

type deleteShowingBiz struct {
	storage DeleteShowingStorage
}

func NewDeleteShowingBiz(storage DeleteShowingStorage) *deleteShowingBiz {
	return &deleteShowingBiz{storage}
}

func (biz *deleteShowingBiz) DeleteShowingById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetShowing(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.ShowingEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.ShowingEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("showing not found"), model.ShowingEntityName)
	}

	if err := biz.storage.DeleteShowing(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.ShowingEntityName)
	}

	return nil
}
