package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCastMember(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.CastMemberCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateCastMemberBiz(storage)

		if err := business.CreateCastMember(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
