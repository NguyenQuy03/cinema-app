package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

type CreateAgeRatingStorage interface {
	CreateAgeRating(ctx context.Context, data *model.AgeRatingCreation) error
}

type createAgeRatingBiz struct {
	storage CreateAgeRatingStorage
}

func NewCreateAgeRatingBiz(storage CreateAgeRatingStorage) *createAgeRatingBiz {
	return &createAgeRatingBiz{storage}
}

func (biz *createAgeRatingBiz) CreateNewAgeRating(ctx context.Context, data *model.AgeRatingCreation) error {

	if err := biz.storage.CreateAgeRating(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.AgeRatingEntityName)
	}

	return nil
}
