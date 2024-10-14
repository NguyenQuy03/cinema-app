package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginSeatTrans "github.com/NguyenQuy03/cinema-app/server/modules/seat/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupSeatRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	seats := v1.Group("seats", middleware.RequireAuth(db))
	{
		seats.POST("", ginSeatTrans.CreateSeat(db))
		seats.GET("/:id", ginSeatTrans.GetSeat(db))
		seats.PATCH("/:id", ginSeatTrans.UpdateSeat(db))
		seats.DELETE("/:id", ginSeatTrans.DeleteSeat(db))
	}
}
