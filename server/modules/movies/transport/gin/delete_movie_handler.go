package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewDeleteMovieBiz(storage)

		if err := business.DeleteMovieById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
