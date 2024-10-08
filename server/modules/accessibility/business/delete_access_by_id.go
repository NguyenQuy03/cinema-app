package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

type DeleteAccessStorage interface {
	GetAccess(ctx context.Context, conds map[string]interface{}) (*model.Accessibility, error)
	DeleteAccess(ctx context.Context, conds map[string]interface{}) error
}

type deleteAccessBiz struct {
	storage DeleteAccessStorage
}

func NewDeleteAccessBiz(storage DeleteAccessStorage) *deleteAccessBiz {
	return &deleteAccessBiz{
		storage: storage,
	}
}

func (biz *deleteAccessBiz) DeleteAccessById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetAccess(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.AccessEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.AccessEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("accessibility not found"), model.AccessEntityName)
	}

	if err := biz.storage.DeleteAccess(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.AccessEntityName)
	}

	return nil
}
