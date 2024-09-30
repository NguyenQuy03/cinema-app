package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupV1Router(router *gin.Engine, db *gorm.DB, redisDB *redis.Client) *gin.Engine {
	v1 := router.Group("v1")
	{
		setupAuthRoutes(v1, db, redisDB)

		setupMovieRoutes(v1, db, redisDB)
		setupGenreRoutes(v1, db, redisDB)
		setupDirectorRoutes(v1, db, redisDB)
	}

	return router
}
