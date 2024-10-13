package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
)

type UpdateCastMemberStorage interface {
	UpdateCastMember(ctx context.Context, conds map[string]interface{}, newData *model.CastMemberUpdate) error
	GetCastMember(ctx context.Context, conds map[string]interface{}) (*model.CastMember, error)
}

type updateCastMemberBiz struct {
	storage UpdateCastMemberStorage
}

func NewUpdateCastMemberBiz(storage UpdateCastMemberStorage) *updateCastMemberBiz {
	return &updateCastMemberBiz{storage}
}

func (biz *updateCastMemberBiz) UpdateCastMember(ctx context.Context, id int, newData *model.CastMemberUpdate) error {
	oldData, err := biz.storage.GetCastMember(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(err, model.CastMemberEntityName)
		}

		return common.ErrCannotUpdateEntity(err, model.CastMemberEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("castMember not found"), model.CastMemberEntityName)
	}

	if err := biz.storage.UpdateCastMember(ctx, map[string]interface{}{
		"id": id,
	}, newData); err != nil {
		return common.ErrCannotUpdateEntity(err, model.CastMemberEntityName)
	}

	return nil
}
