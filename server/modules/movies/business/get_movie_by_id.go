package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

type GetMovieStorage interface {
	GetMovie(ctx context.Context, conds map[string]interface{}) (*model.Movie, error)
}

type getMovieBiz struct {
	storage GetMovieStorage
}

func NewGetMovieBiz(storage GetMovieStorage) *getMovieBiz {
	return &getMovieBiz{storage}
}

func (biz *getMovieBiz) GetMovieById(ctx context.Context, id int) (*model.Movie, error) {
	result, err := biz.storage.GetMovie(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.MOVIE_ENTITY_NAME)
	}

	return result, nil
}
