package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

type UpdateTicketTypeStorage interface {
	UpdateTicketType(ctx context.Context, conds map[string]interface{}, newData *model.TicketTypeUpdate) error
	GetTicketType(ctx context.Context, conds map[string]interface{}) (*model.TicketType, error)
}

type updateTicketTypeBiz struct {
	storage UpdateTicketTypeStorage
}

func NewUpdateTicketTypeBiz(storage UpdateTicketTypeStorage) *updateTicketTypeBiz {
	return &updateTicketTypeBiz{storage}
}

func (biz *updateTicketTypeBiz) UpdateTicketType(ctx context.Context, id int, newData *model.TicketTypeUpdate) error {
	oldData, err := biz.storage.GetTicketType(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.TicketTypeEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.TicketTypeEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("ticket type not found"), model.TicketTypeEntityName)
	}

	if err := biz.storage.UpdateTicketType(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.TicketTypeEntityName)
	}

	return nil
}
