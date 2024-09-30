package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetDirector(ctx context.Context, conds map[string]interface{}) (*model.Director, error) {
	var data model.Director

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

	}

	return &data, nil
}
