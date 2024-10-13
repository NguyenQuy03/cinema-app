package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBooking(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.BookingCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateBookingBiz(storage)

		if err := business.CreateBooking(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
