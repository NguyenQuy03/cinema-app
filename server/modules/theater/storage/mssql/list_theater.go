package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
)

func (s *sqlStorage) ListTheater(ctx context.Context, p *common.Paging, params ...string) ([]model.Theater, error) {
	var result []model.Theater

	if err := s.db.Table(model.Theater{}.TableName()).Count(&p.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.
		Preload("Cinema").
		Preload("Accessibility").
		Preload("Experience").
		Order("id desc").
		Offset((p.Page - 1) * p.Limit).Limit(p.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
