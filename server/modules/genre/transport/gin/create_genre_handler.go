package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateGenre(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.GenreCreation

		slugProvider := new(common.SlugProvider)

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateGenreBiz(storage, slugProvider)

		if err := business.CreateGenre(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.GenreId))
	}
}
