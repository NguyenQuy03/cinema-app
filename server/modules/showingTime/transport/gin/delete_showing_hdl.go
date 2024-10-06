package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/showingTime/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteShowing(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewDeleteShowingBiz(storage)

		if err := business.DeleteShowingById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
