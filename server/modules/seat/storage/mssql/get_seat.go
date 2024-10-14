package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetSeat(ctx context.Context, conds map[string]interface{}) (*model.Seat, error) {
	var data model.Seat

	if err := s.db.Where(conds).Preload("Theater").Preload("SeatType").First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

	}

	return &data, nil
}
