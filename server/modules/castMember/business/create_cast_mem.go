package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
)

type CreateCastMemberStorage interface {
	CreateCastMember(ctx context.Context, data *model.CastMemberCreation) error
}

type createCastMemberBiz struct {
	storage CreateCastMemberStorage
}

func NewCreateCastMemberBiz(storage CreateCastMemberStorage) *createCastMemberBiz {
	return &createCastMemberBiz{
		storage: storage,
	}
}

func (biz *createCastMemberBiz) CreateCastMember(ctx context.Context, data *model.CastMemberCreation) error {
	if err := biz.storage.CreateCastMember(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, model.CastMemberEntityName)
	}

	return nil
}
