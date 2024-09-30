package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginDirectorTrans "github.com/NguyenQuy03/cinema-app/server/modules/director/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func setupDirectorRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	genres := v1.Group("directors", middleware.RequireAuth(db, redisDB))
	{
		genres.POST("", ginDirectorTrans.CreateDirector(db))
		genres.GET("/:id", ginDirectorTrans.GetDirector(db))
		genres.PATCH("/:id", ginDirectorTrans.UpdateGenre(db))
		genres.DELETE("/:id", ginDirectorTrans.DeleteDirector(db))
	}
}
