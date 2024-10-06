package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/place/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePlace(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.PlaceCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		slugProvider := new(common.SlugProvider)

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreatePlaceBiz(storage, slugProvider)

		if err := business.CreatePlace(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
