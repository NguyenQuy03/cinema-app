package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

type ListShowingStorage interface {
	ListShowing(ctx context.Context, paging *common.Paging, params ...string) ([]model.Showing, error)
}

type listShowingBiz struct {
	storage ListShowingStorage
}

func NewListShowingBiz(storage ListShowingStorage) *listShowingBiz {
	return &listShowingBiz{storage}
}

func (biz *listShowingBiz) ListShowing(ctx context.Context, paging *common.Paging) ([]model.Showing, error) {
	result, err := biz.storage.ListShowing(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.ShowingEntityName)
	}

	return result, nil
}
