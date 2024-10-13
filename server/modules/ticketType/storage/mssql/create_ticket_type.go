package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

func (s *sqlStorage) CreateTicketType(ctx context.Context, data *model.TicketTypeCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
