package adminRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginAgeRatingTrans "github.com/NguyenQuy03/cinema-app/server/modules/ageRating/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupAgeRatingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	ticketTypes := v1.Group("age-ratings", middleware.RequireAuth(db))
	{
		ticketTypes.POST("", ginAgeRatingTrans.CreateAgeRating(db))
		ticketTypes.GET("", ginAgeRatingTrans.ListAgeRating(db))
		ticketTypes.GET("/:code", ginAgeRatingTrans.GetAgeRating(db))
		ticketTypes.PATCH("/:code", ginAgeRatingTrans.UpdateAgeRating(db))
		ticketTypes.DELETE("/:code", ginAgeRatingTrans.DeleteAgeRating(db))
	}
}
