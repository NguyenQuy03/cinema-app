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

func CreateAgeRating(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.AgeRatingCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateAgeRatingBiz(storage)

		if err := business.CreateNewAgeRating(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.RatingCode))
	}
}
