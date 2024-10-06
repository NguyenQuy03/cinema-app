package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginTheaterTrans "github.com/NguyenQuy03/cinema-app/server/modules/theater/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func setupTheaterRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	theaters := v1.Group("theaters", middleware.RequireAuth(db, redisDB))
	{
		theaters.POST("", ginTheaterTrans.CreateTheater(db))
		theaters.GET("", ginTheaterTrans.ListTheater(db))
		theaters.GET("/:id", ginTheaterTrans.GetTheater(db))
		theaters.PATCH("/:id", ginTheaterTrans.UpdateTheater(db))
		theaters.DELETE("/:id", ginTheaterTrans.DeleteTheater(db))
	}
}
