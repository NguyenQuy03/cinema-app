package userRoutes

import (
	ginDirectorTrans "github.com/NguyenQuy03/cinema-app/server/modules/director/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupDirectorRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	genres := v1.Group("directors")
	{
		genres.GET("/:id", ginDirectorTrans.GetDirector(db))
	}
}
