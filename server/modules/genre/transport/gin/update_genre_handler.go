package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateGenre(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.GenreUpdate
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
		business := business.NewUpdateGenreBiz(storage)

		if err := business.UpdateGenre(ctx, id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
