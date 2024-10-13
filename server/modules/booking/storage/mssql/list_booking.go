package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
)

func (s *sqlStorage) ListBooking(ctx context.Context, p *common.Paging, params ...string) ([]model.Booking, error) {
	var result []model.Booking

	if err := s.db.Table(model.Booking{}.TableName()).Count(&p.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Order("id desc").Offset((p.Page - 1) * p.Limit).Limit(p.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
