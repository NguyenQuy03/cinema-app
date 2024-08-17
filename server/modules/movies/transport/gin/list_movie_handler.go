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

func ListMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var p common.Paging

		if err := ctx.ShouldBind(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		p.Process()

		var filter model.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewListMovieBiz(storage)

		data, err := business.ListMovie(ctx.Request.Context(), &filter, &p)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.NewAppResponse(data, p, nil))
	}
}
