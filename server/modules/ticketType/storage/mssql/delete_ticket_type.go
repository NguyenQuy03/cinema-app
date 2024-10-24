package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
)

func (s *sqlStorage) DeleteTicketType(ctx context.Context, conds map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.TicketType{}.TableName()).Where(conds).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
