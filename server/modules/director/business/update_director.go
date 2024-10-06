package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
)

type UpdateDirectorStorage interface {
	GetDirector(ctx context.Context, conds map[string]interface{}) (*model.Director, error)
	UpdateDirector(ctx context.Context, conds map[string]interface{}, newData *model.DirectorUpdate) error
}

type updateDirectorBiz struct {
	storage UpdateDirectorStorage
}

func NewUpdateDirectorBiz(storage UpdateDirectorStorage) *updateDirectorBiz {
	return &updateDirectorBiz{storage}
}

func (biz *updateDirectorBiz) UpdateDirector(ctx context.Context, id int, newData *model.DirectorUpdate) error {
	oldData, err := biz.storage.GetDirector(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.DirectorEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.DirectorEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("genre not found"), model.DirectorEntityName)
	}

	if err := biz.storage.UpdateDirector(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.DirectorEntityName)
	}

	return nil
}
