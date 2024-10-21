package userRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginAccessTrans "github.com/NguyenQuy03/cinema-app/server/modules/accessibility/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupAccessRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	ticketTypes := v1.Group("accesses", middleware.RequireAuth(db))
	{
		ticketTypes.GET("", ginAccessTrans.ListAccessibility(db))
	}
}
