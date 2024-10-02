package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
)

type DeleteExperienceStorage interface {
	GetExperience(ctx context.Context, conds map[string]interface{}) (*model.Experience, error)
	DeleteExperience(ctx context.Context, conds map[string]interface{}) error
}

type deleteExperienceBiz struct {
	storage DeleteExperienceStorage
}

func NewDeleteExperienceBiz(storage DeleteExperienceStorage) *deleteExperienceBiz {
	return &deleteExperienceBiz{
		storage: storage,
	}
}

func (biz *deleteExperienceBiz) DeleteExperienceById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetExperience(ctx, map[string]interface{}{"exp_id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.ExperienceEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.ExperienceEntityName)
	}

	if oldData.ExpId == 0 {
		return common.ErrEntityNotFound(errors.New("experience not found"), model.ExperienceEntityName)
	}

	if err := biz.storage.DeleteExperience(ctx, map[string]interface{}{"exp_id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.ExperienceEntityName)
	}

	return nil
}
