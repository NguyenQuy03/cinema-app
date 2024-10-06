package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

type CreateTheaterStorage interface {
	CreateTheater(ctx context.Context, data *model.TheaterCreation) error
}

type createTheaterBiz struct {
	storage CreateTheaterStorage
}

func NewCreateTheaterBiz(storage CreateTheaterStorage) *createTheaterBiz {
	return &createTheaterBiz{storage}
}

func (biz *createTheaterBiz) CreateNewTheater(ctx context.Context, data *model.TheaterCreation) error {

	if err := biz.storage.CreateTheater(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.TheaterEntityName)
	}

	return nil
}
