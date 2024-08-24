package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
)

type ListMovieStorage interface {
	ListMovie(ctx context.Context, filter *model.Filter, paging *common.Paging, params ...string) ([]model.Movie, error)
}

type listMovieBiz struct {
	storage ListMovieStorage
}

func NewListMovieBiz(storage ListMovieStorage) *listMovieBiz {
	return &listMovieBiz{storage}
}

func (biz *listMovieBiz) ListMovie(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.Movie, error) {
	result, err := biz.storage.ListMovie(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.MovieEntityName)
	}

	return result, nil
}
