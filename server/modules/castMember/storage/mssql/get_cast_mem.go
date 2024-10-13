package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetCastMember(ctx context.Context, conds map[string]interface{}) (*model.CastMember, error) {
	var data model.CastMember

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

	}

	return &data, nil
}
