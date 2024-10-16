package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
<<<<<<< HEAD
)

func (s *sqlStorage) CreateBookingSeat(ctx context.Context, data []*model.BookingSeatCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
=======
	"gorm.io/gorm"
)

func (s *sqlStorage) CreateBookingSeat(ctx context.Context, data []*model.BookingSeatCreation, tx *gorm.DB) error {
	if err := tx.Create(data).Error; err != nil {
>>>>>>> development
		return common.ErrDB(err)
	}

	return nil
}
