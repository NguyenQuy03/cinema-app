package mssql

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
)

func (s *sqlStorage) UpdateCastMember(ctx context.Context, conds map[string]interface{}, newData *model.CastMemberUpdate) error {
	if err := s.db.Where(conds).Updates(newData).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
