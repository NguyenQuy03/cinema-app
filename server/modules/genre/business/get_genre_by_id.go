package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

type GetGenreStorage interface {
	GetGenre(ctx context.Context, conds map[string]interface{}) (*model.Genre, error)
}

type getGenreBiz struct {
	storage GetGenreStorage
}

func NewGetGenreBiz(storage GetGenreStorage) *getGenreBiz {
	return &getGenreBiz{storage}
}

func (biz *getGenreBiz) GetGenreById(ctx context.Context, id int) (*model.Genre, error) {
	result, err := biz.storage.GetGenre(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.GenreEntityName)
	}

	return result, nil
}
