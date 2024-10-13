package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
)

func (s *sqlStorage) UpdateMovie(ctx context.Context, conds map[string]interface{}, newData *model.MovieUpdate) error {
	movieId := newData.Id

	// Start a new transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return common.ErrDB(tx.Error)
	}

	// Ensure the transaction is rolled back in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Clear existing genres
	if err := tx.Where("movie_id = ?", movieId).Delete(&model.MovieGenre{}).Error; err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	// Add new genres
	for _, genre := range newData.Genres {
		movieGenre := model.MovieGenre{
			MovieId: movieId,
			GenreId: genre.Id,
		}

		if err := tx.Create(&movieGenre).Error; err != nil {
			tx.Rollback()
			return common.ErrDB(err)
		}
	}
	// Clear existing casts
	if err := tx.Where("movie_id = ?", movieId).Delete(&model.MovieCast{}).Error; err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	// Add new casts
	for _, cast := range newData.CastMembers {
		movieCast := model.MovieCast{
			MovieId: movieId,
			CastId:  cast.Id,
		}

		if err := tx.Create(&movieCast).Error; err != nil {
			tx.Rollback()
			return common.ErrDB(err)
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
