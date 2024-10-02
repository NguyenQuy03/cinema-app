package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
)

func (s *sqlStorage) ListMovie(ctx context.Context, filter *model.Filter, p *common.Paging, params ...string) ([]model.Movie, error) {
	var result []model.Movie

	db := s.db.Where("status <> ?", "inactive")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.Movie{}.TableName()).Count(&p.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Order("movie_id desc").Offset((p.Page - 1) * p.Limit).Limit(p.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
