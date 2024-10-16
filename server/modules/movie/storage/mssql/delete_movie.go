package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
)

func (s *sqlStorage) DeleteMovie(ctx context.Context, conds map[string]interface{}) error {
	inactiveStatus := model.MovieInActiveStatus

	statusValue := inactiveStatus.String()

	if err := s.db.
		Table(model.Movie{}.TableName()).
		Where(conds).
		Updates(map[string]interface{}{"status": statusValue}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
