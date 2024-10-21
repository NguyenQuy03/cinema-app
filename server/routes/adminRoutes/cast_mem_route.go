package adminRoutes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginCastMemTrans "github.com/NguyenQuy03/cinema-app/server/modules/castMember/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupCastMemberRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	castMembers := v1.Group("cast-members", middleware.RequireAuth(db))
	{
		castMembers.POST("", ginCastMemTrans.CreateCastMember(db))
		// castMembers.GET("", ginMovieTrans.ListMovie(db))
		castMembers.GET("/:id", ginCastMemTrans.GetCastMember(db))
		castMembers.PATCH("/:id", ginCastMemTrans.UpdateCastMember(db))
		castMembers.DELETE("/:id", ginCastMemTrans.DeleteCastMember(db))
	}
}
