package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
)

type UpdateExperienceStorage interface {
	UpdateExperience(ctx context.Context, conds map[string]interface{}, newData *model.ExperienceUpdate) error
	GetExperience(ctx context.Context, conds map[string]interface{}) (*model.Experience, error)
}

type updateExperienceBiz struct {
	storage UpdateExperienceStorage
}

func NewUpdateExperienceBiz(storage UpdateExperienceStorage) *updateExperienceBiz {
	return &updateExperienceBiz{storage}
}

func (biz *updateExperienceBiz) UpdateExperience(ctx context.Context, id int, newData *model.ExperienceUpdate) error {
	oldData, err := biz.storage.GetExperience(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.ExperienceEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.ExperienceEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("genre not found"), model.ExperienceEntityName)
	}

	if err := biz.storage.UpdateExperience(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.ExperienceEntityName)
	}

	return nil
}
