package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

type CreateAccessStorage interface {
	CreateAccess(ctx context.Context, data *model.AccessCreation) error
}

type createAccessBiz struct {
	storage CreateAccessStorage
}

func NewCreateAccessBiz(storage CreateAccessStorage) *createAccessBiz {
	return &createAccessBiz{storage}
}

func (biz *createAccessBiz) CreateAccess(ctx context.Context, data *model.AccessCreation) error {

	if err := biz.storage.CreateAccess(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.AccessEntityName)
	}

	return nil
}
