package sqlsv

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/model"
	"gorm.io/gorm"
)

func (s *sqlStorage) GetUser(ctx context.Context, conds map[string]interface{}) (*model.User, error) {
	var data model.User

	if err := s.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
