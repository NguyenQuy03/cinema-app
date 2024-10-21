package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateAgeRating(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.AgeRatingUpdate

		code := ctx.Param("code")

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewUpdateAgeRatingBiz(storage)

		if err := business.UpdateAgeRatingByCode(ctx.Request.Context(), code, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
