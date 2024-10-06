package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
)

func (s *sqlStorage) ListCinema(ctx context.Context, p *common.Paging, params ...string) ([]model.Cinema, error) {
	var result []model.Cinema

	if err := s.db.Table(model.Cinema{}.TableName()).Count(&p.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Order("id desc").Offset((p.Page - 1) * p.Limit).Limit(p.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
