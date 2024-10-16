package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetAccess(ctx context.Context, conds map[string]interface{}) (*model.Accessibility, error) {
	var data model.Accessibility

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil

}
