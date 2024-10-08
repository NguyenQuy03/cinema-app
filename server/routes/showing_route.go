package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginShowingTrans "github.com/NguyenQuy03/cinema-app/server/modules/showingTime/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupShowingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	showings := v1.Group("showings", middleware.RequireAuth(db))
	{
		showings.POST("", ginShowingTrans.CreateShowing(db))
		showings.GET("", ginShowingTrans.ListShowing(db))
		showings.GET("/:id", ginShowingTrans.GetShowing(db))
		showings.PATCH("/:id", ginShowingTrans.UpdateShowing(db))
		showings.DELETE("/:id", ginShowingTrans.DeleteShowing(db))
	}
}
