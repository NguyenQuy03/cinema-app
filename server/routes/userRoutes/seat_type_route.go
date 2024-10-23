package userRoutes

import (
	ginSeatTypeTrans "github.com/NguyenQuy03/cinema-app/server/modules/seatType/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupSeatTypeRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	seatTypes := v1.Group("seat-types")
	{
		seatTypes.GET("", ginSeatTypeTrans.ListSeatType(db))
	}
}
