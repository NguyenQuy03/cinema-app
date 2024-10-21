package commonRoutes

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupCommonV1Router(router *gin.Engine, db *gorm.DB, redisDB *redis.Client, cld *cloudinary.Cloudinary) *gin.Engine {
	v1 := router.Group("v1")
	{
		// Auth
		setupAuthRoutes(v1, db, redisDB)

		// Upload
		setupUploadRoutes(v1, db, cld)
	}

	return router
}
