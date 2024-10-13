package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ticketType/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetTicketType(ctx context.Context, conds map[string]interface{}) (*model.TicketType, error) {
	var data model.TicketType

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

	}

	return &data, nil
}
