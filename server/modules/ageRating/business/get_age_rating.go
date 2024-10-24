package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

type GetAgeRatingStorage interface {
	GetAgeRating(ctx context.Context, conds map[string]interface{}) (*model.AgeRating, error)
}

type getAgeRatingBiz struct {
	storage GetAgeRatingStorage
}

func NewGetAgeRatingBiz(storage GetAgeRatingStorage) *getAgeRatingBiz {
	return &getAgeRatingBiz{storage}
}

func (biz *getAgeRatingBiz) GetAgeRatingByCode(ctx context.Context, code string) (*model.AgeRating, error) {
	result, err := biz.storage.GetAgeRating(ctx, map[string]interface{}{
		"rating_code": code,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.AgeRatingEntityName)
	}

	return result, nil
}
