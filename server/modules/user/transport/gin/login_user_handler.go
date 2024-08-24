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

func AuthenticateUser(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var user model.UserLogin
		var authResponse model.AuthResponse

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewLoginUserBiz(storage)

		if err := business.AuthenticateUser(ctx.Request.Context(), &user, &authResponse); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie("refresh_token", authResponse.RefreshToken, int(model.RefreshTokenMaxAge), "", "", false, true)

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(authResponse))
	}
}
