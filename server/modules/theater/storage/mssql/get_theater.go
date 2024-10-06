package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetTheater(ctx context.Context, conds map[string]interface{}) (*model.Theater, error) {
	var data model.Theater

	if err := s.db.Preload("Cinema").Preload("Accessibility").Preload("Experience").Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
