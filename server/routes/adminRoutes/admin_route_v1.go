package adminRoutes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAdminV1Router(router *gin.Engine, db *gorm.DB) *gin.Engine {
	v1 := router.Group("v1/admin")
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
	}

	return router
}
