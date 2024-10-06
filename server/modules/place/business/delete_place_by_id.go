package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
)

type DeletePlaceStorage interface {
	GetPlace(ctx context.Context, conds map[string]interface{}) (*model.Place, error)
	DeletePlace(ctx context.Context, conds map[string]interface{}) error
}

type deletePlaceBiz struct {
	storage DeletePlaceStorage
}

func NewDeletePlaceBiz(storage DeletePlaceStorage) *deletePlaceBiz {
	return &deletePlaceBiz{
		storage: storage,
	}
}

func (biz *deletePlaceBiz) DeletePlaceById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetPlace(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.PlaceEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.PlaceEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("place not found"), model.PlaceEntityName)
	}

	if err := biz.storage.DeletePlace(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.PlaceEntityName)
	}

	return nil
}
