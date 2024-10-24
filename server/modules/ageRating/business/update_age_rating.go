package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

type UpdateAgeRatingStorage interface {
	GetAgeRating(ctx context.Context, conds map[string]interface{}) (*model.AgeRating, error)
	UpdateAgeRating(ctx context.Context, conds map[string]interface{}, newData *model.AgeRatingUpdate) error
}

type updateAgeRatingBiz struct {
	storage UpdateAgeRatingStorage
}

func NewUpdateAgeRatingBiz(storage UpdateAgeRatingStorage) *updateAgeRatingBiz {
	return &updateAgeRatingBiz{storage}
}

func (biz *updateAgeRatingBiz) UpdateAgeRatingByCode(ctx context.Context, ratingCode string, newData *model.AgeRatingUpdate) error {
	oldData, err := biz.storage.GetAgeRating(ctx, map[string]interface{}{
		"rating_code": ratingCode,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.AgeRatingEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.AgeRatingEntityName)
	}

	if oldData.RatingCode == "" {
		return common.ErrEntityNotFound(errors.New("age rating not found"), model.AgeRatingEntityName)
	}

	if err := biz.storage.UpdateAgeRating(ctx, map[string]interface{}{
		"rating_code": ratingCode,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.AgeRatingEntityName)
	}

	return nil
}
