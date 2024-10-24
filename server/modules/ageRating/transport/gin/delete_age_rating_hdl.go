package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAgeRating(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")

		storage := mssql.NewSQLStorage(db)
		business := business.NewDeleteAgeRatingBiz(storage)

		if err := business.DeleteAgeRatingByCode(ctx.Request.Context(), code); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
