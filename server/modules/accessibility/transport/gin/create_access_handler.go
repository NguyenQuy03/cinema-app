package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAccessibility(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.AccessCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateAccessBiz(storage)

		if err := business.CreateAccess(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
