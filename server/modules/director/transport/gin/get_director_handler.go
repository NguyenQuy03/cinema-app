package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/director/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDirector(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewGetDirectorBiz(storage)

		data, err := business.GetDirectorById(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data))
	}
}
