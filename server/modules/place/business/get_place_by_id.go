package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
)

type GetPlaceStorage interface {
	GetPlace(ctx context.Context, conds map[string]interface{}) (*model.Place, error)
}

type getPlaceBiz struct {
	storage GetPlaceStorage
}

func NewGetPlaceBiz(storage GetPlaceStorage) *getPlaceBiz {
	return &getPlaceBiz{storage}
}

func (biz *getPlaceBiz) GetPlaceById(ctx context.Context, id int) (*model.Place, error) {
	result, err := biz.storage.GetPlace(ctx, map[string]interface{}{
		"place_id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.PlaceEntityName)
	}

	return result, nil
}
