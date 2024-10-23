package adminRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAdminV1Router(router *gin.Engine, db *gorm.DB) *gin.Engine {
	v1 := router.Group("v1/admin", middleware.RequireAuth(db), middleware.RequireAdminRole())
	{

		// Age Rating
		setupAgeRatingRoutes(v1, db)

		// Accessibility
		setupAccessRoutes(v1, db)

		// Cast Member
		setupCastMemberRoutes(v1, db)

		// Genre
		setupGenreRoutes(v1, db)

		// Seat Type
		setupSeatTypeRoutes(v1, db)

		// Cinema
		setupCinemaRoutes(v1, db)

		// Director
		setupDirectorRoutes(v1, db)

		// Experience
		setupExperienceRoutes(v1, db)

		// Movie
		setupMovieRoutes(v1, db)

		// Place
		setupPlaceRoutes(v1, db)

		// Seat
		setupSeatRoutes(v1, db)

		// Showing
		setupShowingRoutes(v1, db)

		// Theater
		setupTheaterRoutes(v1, db)

		// Ticket Type
		setupTicketTypeRoutes(v1, db)
	}

	return router
}
