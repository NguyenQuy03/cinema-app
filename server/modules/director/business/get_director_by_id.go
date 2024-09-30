package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
)

type GetDirectorStorage interface {
	GetDirector(ctx context.Context, conds map[string]interface{}) (*model.Director, error)
}

type getDirectorBiz struct {
	storage GetDirectorStorage
}

func NewGetDirectorBiz(storage GetDirectorStorage) *getDirectorBiz {
	return &getDirectorBiz{storage}
}

func (biz *getDirectorBiz) GetDirectorById(ctx context.Context, id int) (*model.Director, error) {
	result, err := biz.storage.GetDirector(ctx, map[string]interface{}{
		"director_id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.DirectorEntityName)
	}

	return result, nil
}
