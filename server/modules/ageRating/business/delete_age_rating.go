package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

type DeleteAgeRatingStorage interface {
	GetAgeRating(ctx context.Context, conds map[string]interface{}) (*model.AgeRating, error)
	DeleteAgeRating(ctx context.Context, conds map[string]interface{}) error
}

type deleteAgeRatingBiz struct {
	storage DeleteAgeRatingStorage
}

func NewDeleteAgeRatingBiz(storage DeleteAgeRatingStorage) *deleteAgeRatingBiz {
	return &deleteAgeRatingBiz{storage}
}

func (biz *deleteAgeRatingBiz) DeleteAgeRatingByCode(ctx context.Context, code string) error {
	oldData, err := biz.storage.GetAgeRating(ctx, map[string]interface{}{"rating_code": code})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.AgeRatingEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.AgeRatingEntityName)
	}

	if oldData.RatingCode == "" {
		return common.ErrEntityNotFound(errors.New("age rating not found"), model.AgeRatingEntityName)
	}

	if err := biz.storage.DeleteAgeRating(ctx, map[string]interface{}{"rating_code": code}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.AgeRatingEntityName)
	}

	return nil
}
