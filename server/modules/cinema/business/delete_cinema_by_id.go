package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
)

type DeleteCinemaStorage interface {
	GetCinema(ctx context.Context, conds map[string]interface{}) (*model.Cinema, error)
	DeleteCinema(ctx context.Context, conds map[string]interface{}) error
}

type deleteCinemaBiz struct {
	storage DeleteCinemaStorage
}

func NewDeleteCinemaBiz(storage DeleteCinemaStorage) *deleteCinemaBiz {
	return &deleteCinemaBiz{
		storage: storage,
	}
}

func (biz *deleteCinemaBiz) DeleteCinemaById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetCinema(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.CinemaEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.CinemaEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("cinema not found"), model.CinemaEntityName)
	}

	if err := biz.storage.DeleteCinema(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.CinemaEntityName)
	}

	return nil
}
