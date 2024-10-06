package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTheater(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TheaterCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateTheaterBiz(storage)

		if err := business.CreateNewTheater(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
