package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
)

type GetCinemaStorage interface {
	GetCinema(ctx context.Context, conds map[string]interface{}) (*model.Cinema, error)
}

type getCinemaBiz struct {
	storage GetCinemaStorage
}

func NewGetCinemaBiz(storage GetCinemaStorage) *getCinemaBiz {
	return &getCinemaBiz{storage}
}

func (biz *getCinemaBiz) GetCinemaById(ctx context.Context, id int) (*model.Cinema, error) {
	result, err := biz.storage.GetCinema(ctx, map[string]interface{}{
		"cinema_id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.CinemaEntityName)
	}

	return result, nil
}
