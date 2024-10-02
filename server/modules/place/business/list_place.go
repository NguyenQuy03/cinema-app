package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
)

type ListPlaceStorage interface {
	ListPlace(ctx context.Context, p *common.Paging, params ...string) ([]model.Place, error)
}

type listPlaceBiz struct {
	storage ListPlaceStorage
}

func NewListPlaceBiz(storage ListPlaceStorage) *listPlaceBiz {
	return &listPlaceBiz{
		storage: storage,
	}
}

func (biz *listPlaceBiz) ListPlace(ctx context.Context, p *common.Paging, params ...string) ([]model.Place, error) {
	result, err := biz.storage.ListPlace(ctx, p)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.PlaceEntityName)
	}

	return result, nil
}
