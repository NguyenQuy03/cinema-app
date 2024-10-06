package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

type UpdateShowingStorage interface {
	GetShowing(ctx context.Context, conds map[string]interface{}) (*model.Showing, error)
	UpdateShowing(ctx context.Context, conds map[string]interface{}, newData *model.ShowingUpdate) error
}

type updateShowingBiz struct {
	storage UpdateShowingStorage
}

func NewUpdateShowingBiz(storage UpdateShowingStorage) *updateShowingBiz {
	return &updateShowingBiz{storage}
}

func (biz *updateShowingBiz) UpdateShowingById(ctx context.Context, id int, newData *model.ShowingUpdate) error {
	oldData, err := biz.storage.GetShowing(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.ShowingEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.ShowingEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("showing not found"), model.ShowingEntityName)
	}

	if err := biz.storage.UpdateShowing(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.ShowingEntityName)
	}

	return nil
}
