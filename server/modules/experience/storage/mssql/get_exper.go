package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetExperience(ctx context.Context, conds map[string]interface{}) (*model.Experience, error) {
	var data model.Experience

	// Preload genres while fetching the movie
	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
