package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginCinemaTrans "github.com/NguyenQuy03/cinema-app/server/modules/cinema/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupCinemaRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	cinemas := v1.Group("cinemas", middleware.RequireAuth(db))
	{
		cinemas.POST("", ginCinemaTrans.CreateCinema(db))
		cinemas.GET("", ginCinemaTrans.ListCinema(db))
		cinemas.GET("/:id", ginCinemaTrans.GetCinema(db))
		cinemas.PATCH("/:id", ginCinemaTrans.UpdateCinema(db))
		cinemas.DELETE("/:id", ginCinemaTrans.DeleteCinema(db))
	}
}
