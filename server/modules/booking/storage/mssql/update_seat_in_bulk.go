package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateSeatInBulk(ctx context.Context, seatIds []int, field string, value string, tx *gorm.DB) error {
	if err := tx.Model(&model.Seat{}).Where("id IN ?", seatIds).Update(field, value).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
