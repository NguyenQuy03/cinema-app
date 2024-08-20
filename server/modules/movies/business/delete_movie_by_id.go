package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

type DeleteMovieStorage interface {
	GetMovie(ctx context.Context, conds map[string]interface{}) (*model.Movie, error)
	DeleteMovie(ctx context.Context, conds map[string]interface{}) error
}

type deleteMovieBiz struct {
	storage DeleteMovieStorage
}

func NewDeleteMovieBiz(storage DeleteMovieStorage) *deleteMovieBiz {
	return &deleteMovieBiz{storage}
}

func (biz *deleteMovieBiz) DeleteMovieById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetMovie(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(err, model.MOVIE_ENTITY_NAME)
		}

		return common.ErrCannotDeleteEntity(err, model.MOVIE_ENTITY_NAME)
	}

	if *oldData.Status == model.MovieInActiveStatus {
		return model.ErrMovieDeleted
	}

	if err := biz.storage.DeleteMovie(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.MOVIE_ENTITY_NAME)
	}

	return nil
}
