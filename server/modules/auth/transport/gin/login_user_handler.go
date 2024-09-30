package gin

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/mssql"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/redisStorage"
	"github.com/NguyenQuy03/cinema-app/server/utils/cookieUtil"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AuthenticateUser(db *gorm.DB, redisDB *redis.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var user model.UserLogin
		var authResponse *model.AuthResponse

		jwtHandler := new(common.JWTHandler)

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))

			return
		}

		sqlStorage := mssql.NewSQLStorage(db)
		sessionStorage := redisStorage.NewRedisStorage(redisDB)

		business := business.NewLoginUserBiz(sqlStorage, sessionStorage, jwtHandler)

		authResponse, err := business.Login(ctx.Request.Context(), &user)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		cookieUtil.SetCookie(ctx.Writer, common.RefreshToken, authResponse.RefreshToken.Token, authResponse.RefreshToken.ExpiredIn)

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(authResponse))
	}
}
