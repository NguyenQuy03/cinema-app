package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/seatType/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSeatType(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewGetSeatTypeBiz(storage)

		data, err := business.GetSeatTypeById(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data))
	}
}
