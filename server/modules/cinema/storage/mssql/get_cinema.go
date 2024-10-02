package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetCinema(ctx context.Context, conds map[string]interface{}) (*model.Cinema, error) {
	var data model.Cinema

	if err := s.db.Preload("Place").Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
