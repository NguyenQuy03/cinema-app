package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/movie/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.MovieUpdate

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewUpdateMovieBiz(storage)

		if err := business.UpdateMovieById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
