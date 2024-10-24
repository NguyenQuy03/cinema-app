package gintrans

import (
	"errors"
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAgeRating(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		code := ctx.Param("code")

		if code == "" {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(errors.New("bad request")))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewGetAgeRatingBiz(storage)

		data, err := business.GetAgeRatingByCode(ctx.Request.Context(), code)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data))
	}
}
