package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

func (s *sqlStorage) CreateBooking(ctx context.Context, data *model.BookingCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
