package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
)

type GetCastMemberStorage interface {
	GetCastMember(ctx context.Context, conds map[string]interface{}) (*model.CastMember, error)
}

type getCastMemberBiz struct {
	storage GetCastMemberStorage
}

func NewGetCastMemberBiz(storage GetCastMemberStorage) *getCastMemberBiz {
	return &getCastMemberBiz{storage}
}

func (biz *getCastMemberBiz) GetCastMemberById(ctx context.Context, id int) (*model.CastMember, error) {
	result, err := biz.storage.GetCastMember(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(err, model.CastMemberEntityName)
	}

	return result, nil
}
