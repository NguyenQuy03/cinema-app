package userRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginDirectorTrans "github.com/NguyenQuy03/cinema-app/server/modules/director/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupDirectorRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	genres := v1.Group("directors", middleware.RequireAuth(db))
	{
		genres.POST("", ginDirectorTrans.CreateDirector(db))
		genres.GET("/:id", ginDirectorTrans.GetDirector(db))
		genres.PATCH("/:id", ginDirectorTrans.UpdateGenre(db))
		genres.DELETE("/:id", ginDirectorTrans.DeleteDirector(db))
	}
}
