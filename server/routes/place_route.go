package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginPlaceTrans "github.com/NguyenQuy03/cinema-app/server/modules/place/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func setupPlaceRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	places := v1.Group("places", middleware.RequireAuth(db, redisDB))
	{
		places.POST("", ginPlaceTrans.CreatePlace(db))
		places.GET("", ginPlaceTrans.ListPlace(db))
		places.GET("/:id", ginPlaceTrans.GetPlace(db))
		places.PATCH("/:id", ginPlaceTrans.UpdatePlace(db))
		places.DELETE("/:id", ginPlaceTrans.DeletePlace(db))
	}
}
