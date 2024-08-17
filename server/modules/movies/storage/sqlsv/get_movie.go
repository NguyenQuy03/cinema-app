package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
)

func (s *sqlStorage) GetMovie(ctx context.Context, conds map[string]interface{}) (*model.Movie, error) {
	var data model.Movie

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
