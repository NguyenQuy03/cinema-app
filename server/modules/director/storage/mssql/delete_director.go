package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
)

func (s *sqlStorage) DeleteDirector(ctx context.Context, conds map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.Director{}.TableName()).Where(conds).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
