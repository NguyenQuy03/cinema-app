package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

type ListTicketTypeStorage interface {
	ListTicketType(ctx context.Context, paging *common.Paging, params ...string) ([]model.TicketType, error)
}

type listTicketTypeBiz struct {
	storage ListTicketTypeStorage
}

func NewListTicketTypeBiz(storage ListTicketTypeStorage) *listTicketTypeBiz {
	return &listTicketTypeBiz{storage}
}

func (biz *listTicketTypeBiz) ListTicketType(ctx context.Context, paging *common.Paging) ([]model.TicketType, error) {
	result, err := biz.storage.ListTicketType(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(err, model.TicketTypeEntityName)
	}

	return result, nil
}
