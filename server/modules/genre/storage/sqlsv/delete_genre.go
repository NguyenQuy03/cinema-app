package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
)

func (s *sqlStorage) DeleteGenre(ctx context.Context, conds map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.Genre{}.TableName()).Where(conds).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
