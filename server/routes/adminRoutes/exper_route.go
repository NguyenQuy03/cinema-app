package adminRoutes

import (
	ginExperTrans "github.com/NguyenQuy03/cinema-app/server/modules/experience/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupExperienceRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	expers := v1.Group("expers")
	{
		expers.POST("", ginExperTrans.CreateExperience(db))
		expers.GET("/:id", ginExperTrans.GetExperience(db))
		expers.PATCH("/:id", ginExperTrans.UpdateExperience(db))
		expers.DELETE("/:id", ginExperTrans.DeleteExperience(db))
	}
}
