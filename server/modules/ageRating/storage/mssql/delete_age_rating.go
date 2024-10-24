package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
)

func (s *sqlStorage) DeleteAgeRating(ctx context.Context, conds map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.AgeRating{}.TableName()).Where(conds).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
