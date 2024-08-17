package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

func (s *sqlStorage) DeleteMovie(ctx context.Context, conds map[string]interface{}) error {
	if err := s.db.
		Table(model.Movie{}.TableName()).
		Where(conds).
		Updates(map[string]interface{}{"status": model.MovieInActiveStatus}).Error; err != nil {
		return err
	}

	return nil
}
