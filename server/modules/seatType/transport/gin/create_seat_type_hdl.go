package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSeatType(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.SeatTypeCreation

		slugProvider := new(common.SlugProvider)

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateSeatTypeBiz(storage, slugProvider)

		if err := business.CreateSeatType(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
