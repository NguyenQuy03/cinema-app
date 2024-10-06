package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

type UpdateTheaterStorage interface {
	GetTheater(ctx context.Context, conds map[string]interface{}) (*model.Theater, error)
	UpdateTheater(ctx context.Context, conds map[string]interface{}, newData *model.TheaterUpdate) error
}

type updateTheaterBiz struct {
	storage UpdateTheaterStorage
}

func NewUpdateTheaterBiz(storage UpdateTheaterStorage) *updateTheaterBiz {
	return &updateTheaterBiz{storage}
}

func (biz *updateTheaterBiz) UpdateTheaterById(ctx context.Context, id int, newData *model.TheaterUpdate) error {
	oldData, err := biz.storage.GetTheater(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.TheaterEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.TheaterEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("theater not found"), model.TheaterEntityName)
	}

	if err := biz.storage.UpdateTheater(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.TheaterEntityName)
	}

	return nil
}
