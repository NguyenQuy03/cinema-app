package userRoutes

import (
	ginAgeRatingTrans "github.com/NguyenQuy03/cinema-app/server/modules/ageRating/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupAgeRatingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	ticketTypes := v1.Group("age-ratings")
	{
		ticketTypes.GET("", ginAgeRatingTrans.ListAgeRating(db))
	}
}
