package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetGenre(ctx context.Context, conds map[string]interface{}) (*model.Genre, error) {
	var data model.Genre

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

	}

	return &data, nil
}
