package adminRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginAccessTrans "github.com/NguyenQuy03/cinema-app/server/modules/accessibility/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupAccessRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	accesses := v1.Group("accesses", middleware.RequireAuth(db))
	{
		accesses.POST("", ginAccessTrans.CreateAccessibility(db))
		accesses.GET("", ginAccessTrans.ListAccessibility(db))
		accesses.GET("/:id", ginAccessTrans.GetAccessibility(db))
		accesses.PATCH("/:id", ginAccessTrans.UpdateAccessibility(db))
		accesses.DELETE("/:id", ginAccessTrans.DeleteAccessibility(db))
	}
}
