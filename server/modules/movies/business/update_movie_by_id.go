package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

type UpdateMovieStorage interface {
	GetMovie(ctx context.Context, conds map[string]interface{}) (*model.Movie, error)
	UpdateMovie(ctx context.Context, conds map[string]interface{}, newData *model.MovieUpdate) error
}

type updateMovieBiz struct {
	storage UpdateMovieStorage
}

func NewUpdateMovieBiz(storage UpdateMovieStorage) *updateMovieBiz {
	return &updateMovieBiz{storage}
}

func (biz *updateMovieBiz) UpdateMovieById(ctx context.Context, id int, newData *model.MovieUpdate) error {
	oldData, err := biz.storage.GetMovie(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(err, model.MOVIE_ENTITY_NAME)
		}

		return common.ErrCannotUpdateEntity(err, model.MOVIE_ENTITY_NAME)
	}

	if *oldData.Status == model.MovieInActiveStatus {
		return model.ErrMovieDeleted
	}

	if err := biz.storage.UpdateMovie(ctx, map[string]interface{}{"id": id}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.MOVIE_ENTITY_NAME)
	}

	return nil
}
