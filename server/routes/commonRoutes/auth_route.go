package commonRoutes

import (
	ginAuthTrans "github.com/NguyenQuy03/cinema-app/server/modules/auth/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func setupAuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	auth := v1.Group("auth")
	{
		auth.POST("/register", ginAuthTrans.RegisterUser(db))
		auth.POST("/login", ginAuthTrans.AuthenticateUser(db, redisDB))
		auth.POST("/refresh-token", ginAuthTrans.RefreshToken(db, redisDB))
	}
}
