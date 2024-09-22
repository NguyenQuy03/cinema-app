package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

type DeleteGenreStorage interface {
	GetGenre(ctx context.Context, conds map[string]interface{}) (*model.Genre, error)
	DeleteGenre(ctx context.Context, conds map[string]interface{}) error
}

type deleteGenreBiz struct {
	storage DeleteGenreStorage
}

func NewDeleteGenreBiz(storage DeleteGenreStorage) *deleteGenreBiz {
	return &deleteGenreBiz{
		storage: storage,
	}
}

func (biz *deleteGenreBiz) DeleteGenreById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetGenre(ctx, map[string]interface{}{"genre_id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.GenreEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.GenreEntityName)
	}

	if oldData.GenreId == 0 {
		return common.ErrEntityNotFound(errors.New("genre not found"), model.GenreEntityName)
	}

	if err := biz.storage.DeleteGenre(ctx, map[string]interface{}{"genre_id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.GenreEntityName)
	}

	return nil
}
