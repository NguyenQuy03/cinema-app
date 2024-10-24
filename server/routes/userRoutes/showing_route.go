package userRoutes

import (
	ginShowingTrans "github.com/NguyenQuy03/cinema-app/server/modules/showingTime/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupShowingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	showings := v1.Group("showings")
	{
		showings.GET("", ginShowingTrans.ListShowing(db))
		showings.GET("/:id", ginShowingTrans.GetShowing(db))
	}
}
