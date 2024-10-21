package commonRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginUploadTrans "github.com/NguyenQuy03/cinema-app/server/modules/upload/transport/gin"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupUploadRoutes(v1 *gin.RouterGroup, db *gorm.DB, cld *cloudinary.Cloudinary) {
	auth := v1.Group("upload", middleware.RequireAuth(db))
	{
		auth.POST("/image", ginUploadTrans.UploadMovieImage(cld))
	}
}
