package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

type GetShowingStorage interface {
	GetShowing(ctx context.Context, conds map[string]interface{}) (*model.Showing, error)
}

type getShowingBiz struct {
	storage GetShowingStorage
}

func NewGetShowingBiz(storage GetShowingStorage) *getShowingBiz {
	return &getShowingBiz{storage}
}

func (biz *getShowingBiz) GetShowingById(ctx context.Context, id int) (*model.Showing, error) {
	result, err := biz.storage.GetShowing(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.ShowingEntityName)
	}

	return result, nil
}
