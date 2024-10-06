package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
)

type UpdateCinemaStorage interface {
	UpdateCinema(ctx context.Context, conds map[string]interface{}, newData *model.CinemaUpdate) error
	GetCinema(ctx context.Context, conds map[string]interface{}) (*model.Cinema, error)
}

type updateCinemaBiz struct {
	storage UpdateCinemaStorage
}

func NewUpdateCinemaBiz(storage UpdateCinemaStorage) *updateCinemaBiz {
	return &updateCinemaBiz{storage}
}

func (biz *updateCinemaBiz) UpdateCinema(ctx context.Context, id int, newData *model.CinemaUpdate) error {
	oldData, err := biz.storage.GetCinema(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.CinemaEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.CinemaEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("cinema not found"), model.CinemaEntityName)
	}

	if err := biz.storage.UpdateCinema(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.CinemaEntityName)
	}

	return nil
}
