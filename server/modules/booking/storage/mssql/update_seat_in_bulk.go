package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

func (s *sqlStorage) UpdateSeatInBulk(ctx context.Context, seatIds []int, field string, value string) error {
	if err := s.db.Model(&model.BookingSeat{}).Where("id IN ?", seatIds).Update(field, value).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
