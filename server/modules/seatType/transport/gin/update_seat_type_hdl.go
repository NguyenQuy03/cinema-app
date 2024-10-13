package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateSeatType(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.SeatTypeUpdate
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
		business := business.NewUpdateSeatTypeBiz(storage)

		if err := business.UpdateSeatType(ctx, id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
