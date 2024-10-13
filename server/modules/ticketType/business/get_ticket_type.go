package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

type GetTicketTypeStorage interface {
	GetTicketType(ctx context.Context, conds map[string]interface{}) (*model.TicketType, error)
}

type getTicketTypeBiz struct {
	storage GetTicketTypeStorage
}

func NewGetTicketTypeBiz(storage GetTicketTypeStorage) *getTicketTypeBiz {
	return &getTicketTypeBiz{storage}
}

func (biz *getTicketTypeBiz) GetTicketTypeById(ctx context.Context, id int) (*model.TicketType, error) {
	result, err := biz.storage.GetTicketType(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.TicketTypeEntityName)
	}

	return result, nil
}
