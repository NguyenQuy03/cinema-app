package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

func (s *sqlStorage) CreateMovie(ctx context.Context, data *model.MovieCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
