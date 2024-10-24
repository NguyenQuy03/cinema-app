package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

func (s *sqlStorage) ListAgeRating(ctx context.Context, p *common.Paging, params ...string) ([]model.AgeRating, error) {
	var result []model.AgeRating

	if err := s.db.Table(model.AgeRating{}.TableName()).Count(&p.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.
		Order("rating_value desc").
		Offset((p.Page - 1) * p.Limit).Limit(p.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
