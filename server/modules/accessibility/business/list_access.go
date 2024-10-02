package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

type ListAccessStorage interface {
	ListAccess(ctx context.Context, p *common.Paging, params ...string) ([]model.Accessibility, error)
}

type listAccessBiz struct {
	storage ListAccessStorage
}

func NewListAccessBiz(storage ListAccessStorage) *listAccessBiz {
	return &listAccessBiz{
		storage: storage,
	}
}

func (biz *listAccessBiz) ListAccess(ctx context.Context, p *common.Paging, params ...string) ([]model.Accessibility, error) {
	result, err := biz.storage.ListAccess(ctx, p)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.AccessEntityName)
	}

	return result, nil
}
