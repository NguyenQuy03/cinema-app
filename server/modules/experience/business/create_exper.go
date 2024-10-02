package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
)

type CreateExperienceStorage interface {
	CreateExperience(ctx context.Context, data *model.ExperienceCreation) error
}

type createExperienceBiz struct {
	storage CreateExperienceStorage
}

func NewCreateExperienceBiz(storage CreateExperienceStorage) *createExperienceBiz {
	return &createExperienceBiz{storage}
}

func (biz *createExperienceBiz) CreateExperience(ctx context.Context, data *model.ExperienceCreation) error {

	if err := biz.storage.CreateExperience(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.ExperienceEntityName)
	}

	return nil
}
