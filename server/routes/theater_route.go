package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginTheaterTrans "github.com/NguyenQuy03/cinema-app/server/modules/theater/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupTheaterRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	theaters := v1.Group("theaters", middleware.RequireAuth(db))
	{
		theaters.POST("", ginTheaterTrans.CreateTheater(db))
		theaters.GET("", ginTheaterTrans.ListTheater(db))
		theaters.GET("/:id", ginTheaterTrans.GetTheater(db))
		theaters.PATCH("/:id", ginTheaterTrans.UpdateTheater(db))
		theaters.DELETE("/:id", ginTheaterTrans.DeleteTheater(db))
	}
}
