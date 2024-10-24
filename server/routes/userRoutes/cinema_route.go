package userRoutes

import (
	ginCinemaTrans "github.com/NguyenQuy03/cinema-app/server/modules/cinema/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupCinemaRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	cinemas := v1.Group("cinemas")
	{
		cinemas.GET("", ginCinemaTrans.ListCinema(db))
		cinemas.GET("/:id", ginCinemaTrans.GetCinema(db))
	}
}
