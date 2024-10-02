package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

func (s *sqlStorage) ListAccess(ctx context.Context, p *common.Paging, params ...string) ([]model.Accessibility, error) {
	var result []model.Accessibility

	if err := s.db.Table(model.Accessibility{}.TableName()).Count(&p.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Order("acc_id desc").Offset((p.Page - 1) * p.Limit).Limit(p.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
