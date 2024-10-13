package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
)

func (s *sqlStorage) CreateCastMember(ctx context.Context, data *model.CastMemberCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
