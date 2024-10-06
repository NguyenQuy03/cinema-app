package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

type ListTheaterStorage interface {
	ListTheater(ctx context.Context, paging *common.Paging, params ...string) ([]model.Theater, error)
}

type listTheaterBiz struct {
	storage ListTheaterStorage
}

func NewListTheaterBiz(storage ListTheaterStorage) *listTheaterBiz {
	return &listTheaterBiz{storage}
}

func (biz *listTheaterBiz) ListTheater(ctx context.Context, paging *common.Paging) ([]model.Theater, error) {
	result, err := biz.storage.ListTheater(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.TheaterEntityName)
	}

	return result, nil
}
