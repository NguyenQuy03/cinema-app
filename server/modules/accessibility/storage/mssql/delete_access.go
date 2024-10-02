package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
)

func (s *sqlStorage) DeleteAccess(ctx context.Context, conds map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.Accessibility{}.TableName()).Where(conds).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
