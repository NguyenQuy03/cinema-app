package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/movies/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.MovieCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewCreateMovieBiz(storage)

		if err := business.CreateNewMovie(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
