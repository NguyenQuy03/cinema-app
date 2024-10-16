package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginMovieTrans "github.com/NguyenQuy03/cinema-app/server/modules/movie/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupMovieRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	movies := v1.Group("movies", middleware.RequireAuth(db))
	{
		movies.POST("", ginMovieTrans.CreateMovie(db))
		movies.GET("", ginMovieTrans.ListMovie(db))
		movies.GET("/:id", ginMovieTrans.GetMovie(db))
		movies.PATCH("/:id", ginMovieTrans.UpdateMovie(db))
		movies.DELETE("/:id", ginMovieTrans.DeleteMovie(db))
	}
}
