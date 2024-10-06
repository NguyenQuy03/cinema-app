package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
)

type CreateShowingStorage interface {
	CreateShowing(ctx context.Context, data *model.ShowingCreation) error
}

type createShowingBiz struct {
	storage CreateShowingStorage
}

func NewCreateShowingBiz(storage CreateShowingStorage) *createShowingBiz {
	return &createShowingBiz{storage}
}

func (biz *createShowingBiz) CreateNewShowing(ctx context.Context, data *model.ShowingCreation) error {
	if err := biz.storage.CreateShowing(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.ShowingEntityName)
	}

	return nil
}
