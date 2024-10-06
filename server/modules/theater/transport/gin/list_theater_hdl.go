package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/theater/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTheater(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var p common.Paging

		if err := ctx.ShouldBind(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		p.Process()

		storage := mssql.NewSQLStorage(db)
		business := business.NewListTheaterBiz(storage)

		data, err := business.ListTheater(ctx.Request.Context(), &p)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewAppResponse(data, p, nil))
	}
}
