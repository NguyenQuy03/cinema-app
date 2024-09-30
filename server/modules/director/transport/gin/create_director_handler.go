package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDirector(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.DirectorCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateDirectorBiz(storage)

		if err := business.CreateDirector(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.DirectorId))
	}
}
