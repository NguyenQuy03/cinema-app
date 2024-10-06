package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateShowing(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.ShowingCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateShowingBiz(storage)

		if err := business.CreateNewShowing(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
