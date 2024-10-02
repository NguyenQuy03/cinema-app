package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/cinema/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCinema(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.CinemaCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		slugProvider := new(common.SlugProvider)

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateCinemaBiz(storage, slugProvider)

		if err := business.CreateCinema(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.CinemaId))
	}
}
