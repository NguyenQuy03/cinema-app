package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
)

type GetUserStorage interface {
	GetUser(ctx context.Context, conds map[string]interface{}) (*model.User, error)
}

type getUserBiz struct {
	storage GetUserStorage
}

func NewGetUserBiz(storage GetUserStorage) *getUserBiz {
	return &getUserBiz{storage: storage}
}

func (biz *getUserBiz) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	// Find User by email
	user, err := biz.storage.GetUser(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.UserEntityName)
	}

	return user, nil
}
