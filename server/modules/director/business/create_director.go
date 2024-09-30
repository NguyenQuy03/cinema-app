package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
)

type CreateDirectorStorage interface {
	CreateDirector(ctx context.Context, data *model.DirectorCreation) error
}

type createDirectorBiz struct {
	storage CreateDirectorStorage
}

func NewCreateDirectorBiz(storage CreateDirectorStorage) *createDirectorBiz {
	return &createDirectorBiz{storage}
}

func (biz *createDirectorBiz) CreateDirector(ctx context.Context, data *model.DirectorCreation) error {
	// name := strings.TrimSpace(data.DirectorName)

	// if name == "" {
	// 	return model.ErrDirectorNameIsBlank
	// }

	if err := biz.storage.CreateDirector(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.DirectorEntityName)
	}

	return nil
}
