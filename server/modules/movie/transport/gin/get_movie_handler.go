package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewGetMovieBiz(storage)

		data, err := business.GetMovieById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data))
	}
}
