package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupV1Router(router *gin.Engine, db *gorm.DB, redisDB *redis.Client) *gin.Engine {
	v1 := router.Group("v1")
	{
		// Auth
		setupAuthRoutes(v1, db, redisDB)

		// Movie
		setupMovieRoutes(v1, db, redisDB)

		// Genre
		setupGenreRoutes(v1, db, redisDB)

		// Director
		setupDirectorRoutes(v1, db, redisDB)

		// Experience
		setupExperienceRoutes(v1, db, redisDB)

		// Accessibility
		setupAccessRoutes(v1, db, redisDB)

		// Place
		setupPlaceRoutes(v1, db, redisDB)

		// Cinema
		setupCinemaRoutes(v1, db, redisDB)

		// Theater
		setupTheaterRoutes(v1, db, redisDB)

		// Showing
		setupShowingRoutes(v1, db, redisDB)
	}

	return router
}
