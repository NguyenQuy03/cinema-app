package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/castMember/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCastMember(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.CastMemberUpdate
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewUpdateCastMemberBiz(storage)

		if err := business.UpdateCastMember(ctx, id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
