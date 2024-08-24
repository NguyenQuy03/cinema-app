package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
)

func (s *sqlStorage) UpdateMovie(ctx context.Context, conds map[string]interface{}, newData *model.MovieUpdate) error {
	if err := s.db.Where(conds).Updates(newData).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
