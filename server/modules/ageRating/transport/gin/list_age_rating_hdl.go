package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/ageRating/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAgeRating(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var p common.Paging

		if err := ctx.ShouldBind(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		p.Process()

		storage := mssql.NewSQLStorage(db)
		business := business.NewListAgeRatingBiz(storage)

		data, err := business.ListAgeRating(ctx.Request.Context(), &p)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewAppResponse(data, p, nil))
	}
}
