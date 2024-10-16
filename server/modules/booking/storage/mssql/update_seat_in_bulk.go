package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
<<<<<<< HEAD
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

func (s *sqlStorage) UpdateSeatInBulk(ctx context.Context, seatIds []int, field string, value string) error {
	if err := s.db.Model(&model.BookingSeat{}).Where("id IN ?", seatIds).Update(field, value).Error; err != nil {
=======
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateSeatInBulk(ctx context.Context, seatIds []int, field string, value string, tx *gorm.DB) error {
	if err := tx.Model(&model.Seat{}).Where("id IN ?", seatIds).Update(field, value).Error; err != nil {
>>>>>>> development
		return common.ErrDB(err)
	}
	return nil
}
