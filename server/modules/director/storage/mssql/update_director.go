package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
)

func (s *sqlStorage) UpdateDirector(ctx context.Context, conds map[string]interface{}, newData *model.DirectorUpdate) error {
	if err := s.db.Where(conds).Updates(newData).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
