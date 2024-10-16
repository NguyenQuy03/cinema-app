package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) CreateBooking(ctx context.Context, data *model.BookingCreation, tx *gorm.DB) error {
	if err := tx.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
