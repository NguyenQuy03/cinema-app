package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
)

type UpdatePlaceStorage interface {
	UpdatePlace(ctx context.Context, conds map[string]interface{}, newData *model.PlaceUpdate) error
	GetPlace(ctx context.Context, conds map[string]interface{}) (*model.Place, error)
}

type updatePlaceBiz struct {
	storage UpdatePlaceStorage
}

func NewUpdatePlaceBiz(storage UpdatePlaceStorage) *updatePlaceBiz {
	return &updatePlaceBiz{storage}
}

func (biz *updatePlaceBiz) UpdatePlace(ctx context.Context, id int, newData *model.PlaceUpdate) error {
	oldData, err := biz.storage.GetPlace(ctx, map[string]interface{}{
		"place_id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.PlaceEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.PlaceEntityName)
	}

	if oldData.PlaceId == 0 {
		return common.ErrEntityNotFound(errors.New("place not found"), model.PlaceEntityName)
	}

	if err := biz.storage.UpdatePlace(ctx, map[string]interface{}{
		"place_id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.PlaceEntityName)
	}

	return nil
}
