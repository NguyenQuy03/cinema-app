package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

type GetTheaterStorage interface {
	GetTheater(ctx context.Context, conds map[string]interface{}) (*model.Theater, error)
}

type getTheaterBiz struct {
	storage GetTheaterStorage
}

func NewGetTheaterBiz(storage GetTheaterStorage) *getTheaterBiz {
	return &getTheaterBiz{storage}
}

func (biz *getTheaterBiz) GetTheaterById(ctx context.Context, id int) (*model.Theater, error) {
	result, err := biz.storage.GetTheater(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.TheaterEntityName)
	}

	return result, nil
}
