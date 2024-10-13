package business

import (
	"context"
	"errors"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
)

type DeleteCastMemberStorage interface {
	GetCastMember(ctx context.Context, conds map[string]interface{}) (*model.CastMember, error)
	DeleteCastMember(ctx context.Context, conds map[string]interface{}) error
}

type deleteCastMemberBiz struct {
	storage DeleteCastMemberStorage
}

func NewDeleteCastMemberBiz(storage DeleteCastMemberStorage) *deleteCastMemberBiz {
	return &deleteCastMemberBiz{
		storage: storage,
	}
}

func (biz *deleteCastMemberBiz) DeleteCastMemberById(ctx context.Context, id int) error {
	oldData, err := biz.storage.GetCastMember(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrEntityNotFound(err, model.CastMemberEntityName)
		}

		return common.ErrCannotDeleteEntity(err, model.CastMemberEntityName)
	}

	if oldData.Id == 0 {
		return common.ErrEntityNotFound(errors.New("castMember not found"), model.CastMemberEntityName)
	}

	if err := biz.storage.DeleteCastMember(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(err, model.CastMemberEntityName)
	}

	return nil
}
