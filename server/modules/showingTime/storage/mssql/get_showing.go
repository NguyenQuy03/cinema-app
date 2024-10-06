package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetShowing(ctx context.Context, conds map[string]interface{}) (*model.Showing, error) {
	var data model.Showing

	if err := s.db.Preload("Movie").Preload("Theater").Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
