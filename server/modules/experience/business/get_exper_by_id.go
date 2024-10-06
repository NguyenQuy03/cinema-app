package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
)

type GetExperienceStorage interface {
	GetExperience(ctx context.Context, conds map[string]interface{}) (*model.Experience, error)
}

type getExperienceBiz struct {
	storage GetExperienceStorage
}

func NewGetExperienceBiz(storage GetExperienceStorage) *getExperienceBiz {
	return &getExperienceBiz{storage}
}

func (biz *getExperienceBiz) GetExperienceById(ctx context.Context, id int) (*model.Experience, error) {
	result, err := biz.storage.GetExperience(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.ExperienceEntityName)
	}

	return result, nil
}
