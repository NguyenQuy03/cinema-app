package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

type CreateTicketTypeStorage interface {
	CreateTicketType(ctx context.Context, data *model.TicketTypeCreation) error
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createTicketTypeBiz struct {
	storage      CreateTicketTypeStorage
	slugProvider SlugProvider
}

func NewCreateTicketTypeBiz(storage CreateTicketTypeStorage, slugProvider SlugProvider) *createTicketTypeBiz {
	return &createTicketTypeBiz{
		storage:      storage,
		slugProvider: slugProvider,
	}
}

func (biz *createTicketTypeBiz) CreateTicketType(ctx context.Context, data *model.TicketTypeCreation) error {

	data.Slug = biz.slugProvider.GenerateSlug(data.TicketName)

	if err := biz.storage.CreateTicketType(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.TicketTypeEntityName)
	}

	return nil
}
