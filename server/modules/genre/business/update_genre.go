package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

type UpdateGenreStorage interface {
	UpdateGenre(ctx context.Context, conds map[string]interface{}, newData *model.GenreUpdate) error
	GetGenre(ctx context.Context, conds map[string]interface{}) (*model.Genre, error)
}

type updateGenreBiz struct {
	storage UpdateGenreStorage
}

func NewUpdateGenreBiz(storage UpdateGenreStorage) *updateGenreBiz {
	return &updateGenreBiz{storage}
}

func (biz *updateGenreBiz) UpdateGenre(ctx context.Context, id int, newData *model.GenreUpdate) error {
	oldData, err := biz.storage.GetGenre(ctx, map[string]interface{}{
		"genre_id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.GenreEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.GenreEntityName)
	}

	if oldData.GenreId == 0 {
		return common.ErrEntityNotFound(errors.New("genre not found"), model.GenreEntityName)
	}

	if err := biz.storage.UpdateGenre(ctx, map[string]interface{}{
		"genre_id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.GenreEntityName)
	}

	return nil
}
