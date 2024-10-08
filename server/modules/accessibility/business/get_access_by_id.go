package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

type GetAccessStorage interface {
	GetAccess(ctx context.Context, conds map[string]interface{}) (*model.Accessibility, error)
}

type getAccessBiz struct {
	storage GetAccessStorage
}

func NewGetAccessBiz(storage GetAccessStorage) *getAccessBiz {
	return &getAccessBiz{storage}
}

func (biz *getAccessBiz) GetAccessById(ctx context.Context, id int) (*model.Accessibility, error) {
	result, err := biz.storage.GetAccess(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.AccessEntityName)
	}

	return result, nil
}
