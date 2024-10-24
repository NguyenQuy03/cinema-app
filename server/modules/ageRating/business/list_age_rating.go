package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

type ListAgeRatingStorage interface {
	ListAgeRating(ctx context.Context, paging *common.Paging, params ...string) ([]model.AgeRating, error)
}

type listAgeRatingBiz struct {
	storage ListAgeRatingStorage
}

func NewListAgeRatingBiz(storage ListAgeRatingStorage) *listAgeRatingBiz {
	return &listAgeRatingBiz{storage}
}

func (biz *listAgeRatingBiz) ListAgeRating(ctx context.Context, paging *common.Paging) ([]model.AgeRating, error) {
	result, err := biz.storage.ListAgeRating(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.AgeRatingEntityName)
	}

	return result, nil
}
