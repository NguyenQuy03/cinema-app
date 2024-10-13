package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

type DeleteTicketTypeStorage interface {
	GetTicketType(ctx context.Context, conds map[string]interface{}) (*model.TicketType, error)
	DeleteTicketType(ctx context.Context, conds map[string]interface{}) error
}

type deleteTicketTypeBiz struct {
	storage DeleteTicketTypeStorage
}

func NewDeleteTicketTypeBiz(storage DeleteTicketTypeStorage) *deleteTicketTypeBiz {
	return &deleteTicketTypeBiz{
		storage: storage,
	}
}

func (biz *deleteTicketTypeBiz) DeleteTicketTypeById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetTicketType(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.TicketTypeEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.TicketTypeEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("ticket type not found"), model.TicketTypeEntityName)
	}

	if err := biz.storage.DeleteTicketType(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.TicketTypeEntityName)
	}

	return nil
}
