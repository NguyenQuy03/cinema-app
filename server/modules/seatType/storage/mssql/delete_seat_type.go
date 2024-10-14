package mssql

import (
	"context"
	"time"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
)

func (s *sqlStorage) DeleteSeatType(ctx context.Context, conds map[string]interface{}) error {
	db := s.db

	if err := db.Table(model.SeatType{}.
		TableName()).Where(conds).
		Update("deleted_at", time.Now()).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
