package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

type UpdateAccessStorage interface {
	UpdateAccess(ctx context.Context, conds map[string]interface{}, newData *model.AccessUpdate) error
	GetAccess(ctx context.Context, conds map[string]interface{}) (*model.Accessibility, error)
}

type updateAccessBiz struct {
	storage UpdateAccessStorage
}

func NewUpdateAccessBiz(storage UpdateAccessStorage) *updateAccessBiz {
	return &updateAccessBiz{storage}
}

func (biz *updateAccessBiz) UpdateAccess(ctx context.Context, id int, newData *model.AccessUpdate) error {
	oldData, err := biz.storage.GetAccess(ctx, map[string]interface{}{
		"acc_id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.AccessEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.AccessEntityName)
	}

	if oldData.AccId == 0 {
		return common.ErrEntityNotFound(errors.New("accessibility not found"), model.AccessEntityName)
	}

	if err := biz.storage.UpdateAccess(ctx, map[string]interface{}{
		"acc_id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.AccessEntityName)
	}

	return nil
}
