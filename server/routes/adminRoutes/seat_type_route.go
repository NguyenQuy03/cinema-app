package adminRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginSeatTypeTrans "github.com/NguyenQuy03/cinema-app/server/modules/seatType/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupSeatTypeRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	seatTypes := v1.Group("seat-types", middleware.RequireAuth(db))
	{
		seatTypes.POST("", ginSeatTypeTrans.CreateSeatType(db))
		seatTypes.GET("/:id", ginSeatTypeTrans.GetSeatType(db))
		seatTypes.GET("", ginSeatTypeTrans.ListSeatType(db))
		seatTypes.PATCH("/:id", ginSeatTypeTrans.UpdateSeatType(db))
		seatTypes.DELETE("/:id", ginSeatTypeTrans.DeleteSeatType(db))
	}
}
