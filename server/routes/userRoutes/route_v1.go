package userRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupV1Router(router *gin.Engine, db *gorm.DB, redisDB *redis.Client) *gin.Engine {
	v1 := router.Group("v1", middleware.RequireAuth(db))
	{
		// Movie
		setupMovieRoutes(v1, db)

		// Director
		setupDirectorRoutes(v1, db)

		// Experience
		setupExperienceRoutes(v1, db)

		// Accessibility
		setupAccessRoutes(v1, db)

		// Place
		setupPlaceRoutes(v1, db)

		// Cinema
		setupCinemaRoutes(v1, db)

		// Theater
		setupTheaterRoutes(v1, db)

		// Showing
		setupShowingRoutes(v1, db)

		// Ticket Type
		setupTicketTypeRoutes(v1, db)

		// Seat Type
		setupSeatTypeRoutes(v1, db)

		// Seat
		setupSeatRoutes(v1, db)

		// Booking
		setupBookingRoutes(v1, db)

		// Age Rating
		setupAgeRatingRoutes(v1, db)
	}

	return router
}
