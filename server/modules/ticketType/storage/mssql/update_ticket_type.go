package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

func (s *sqlStorage) UpdateTicketType(ctx context.Context, conds map[string]interface{}, newData *model.TicketTypeUpdate) error {
	if err := s.db.Where(conds).Updates(newData).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}