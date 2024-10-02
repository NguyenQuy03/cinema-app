package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
)

type ListCinemaStorage interface {
	ListCinema(ctx context.Context, p *common.Paging, params ...string) ([]model.Cinema, error)
}

type listCinemaBiz struct {
	storage ListCinemaStorage
}

func NewListCinemaBiz(storage ListCinemaStorage) *listCinemaBiz {
	return &listCinemaBiz{
		storage: storage,
	}
}

func (biz *listCinemaBiz) ListCinema(ctx context.Context, p *common.Paging, params ...string) ([]model.Cinema, error) {
	result, err := biz.storage.ListCinema(ctx, p)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.CinemaEntityName)
	}

	return result, nil
}
