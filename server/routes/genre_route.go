package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginGenreTrans "github.com/NguyenQuy03/cinema-app/server/modules/genre/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupGenreRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	genres := v1.Group("genres", middleware.RequireAuth(db))
	{
		genres.POST("", ginGenreTrans.CreateGenre(db))
		// genres.GET("", ginMovieTrans.ListMovie(db))
		genres.GET("/:id", ginGenreTrans.GetGenre(db))
		genres.PATCH("/:id", ginGenreTrans.UpdateGenre(db))
		genres.DELETE("/:id", ginGenreTrans.DeleteGenre(db))
	}
}
