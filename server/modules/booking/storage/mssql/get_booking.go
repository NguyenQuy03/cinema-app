package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetBooking(ctx context.Context, conds map[string]interface{}) (*model.Booking, error) {
	var data model.Booking

	if err := s.db.
		Where(conds).
		Preload("User").
		Preload("Showing").
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
