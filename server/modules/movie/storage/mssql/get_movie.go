package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetMovie(ctx context.Context, conds map[string]interface{}) (*model.Movie, error) {
	var data model.Movie

	if err := s.db.Preload("Genres").Preload("Director").Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
