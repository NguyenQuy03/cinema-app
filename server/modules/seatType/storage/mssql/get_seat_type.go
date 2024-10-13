package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetSeatType(ctx context.Context, conds map[string]interface{}) (*model.SeatType, error) {
	var data model.SeatType

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

	}

	return &data, nil
}
