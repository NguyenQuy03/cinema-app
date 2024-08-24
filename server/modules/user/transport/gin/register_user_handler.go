package gin

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.UserRegister

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewRegisterUserBiz(storage)

		if err := business.RegisterUser(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
