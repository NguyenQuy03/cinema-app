package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

func (s *sqlStorage) UpdateMovie(ctx context.Context, conds map[string]interface{}, newData *model.MovieUpdate) error {
	if err := s.db.Where(conds).Updates(newData).Error; err != nil {
		return err
	}

	return nil
}
