package gin

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/redisStorage"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RefreshToken(db *gorm.DB, redisDB *redis.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var authResponse model.AuthResponse

		sqlStorage := sqlsv.NewSQLStorage(db)
		sessionStorage := redisStorage.NewRedisStorage(redisDB)

		business := business.NewRefreshTokenBiz(sqlStorage, sessionStorage)

		if err := business.RefreshToken(ctx.Request.Context(), ctx.Request, &authResponse); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(authResponse))
	}
}
