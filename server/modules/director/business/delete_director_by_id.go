package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
)

type DeleteDirectorStorage interface {
	GetDirector(ctx context.Context, conds map[string]interface{}) (*model.Director, error)
	DeleteDirector(ctx context.Context, conds map[string]interface{}) error
}

type deleteDirectorBiz struct {
	storage DeleteDirectorStorage
}

func NewDeleteDirectorBiz(storage DeleteDirectorStorage) *deleteDirectorBiz {
	return &deleteDirectorBiz{
		storage: storage,
	}
}

func (biz *deleteDirectorBiz) DeleteDirectorById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetDirector(ctx, map[string]interface{}{"director_id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.DirectorEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.DirectorEntityName)
	}

	if oldData.DirectorId == 0 {
		return model.ErrDirectorNotFound
	}

	if err := biz.storage.DeleteDirector(ctx, map[string]interface{}{"director_id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.DirectorEntityName)
	}

	return nil
}
