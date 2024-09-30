package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginMovieTrans "github.com/NguyenQuy03/cinema-app/server/modules/movie/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func setupMovieRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	movies := v1.Group("movies", middleware.RequireAuth(db, redisDB))
	{
		movies.POST("", ginMovieTrans.CreateMovie(db))
		movies.GET("", ginMovieTrans.ListMovie(db))
		movies.GET("/:id", ginMovieTrans.GetMovie(db))
		movies.PATCH("/:id", ginMovieTrans.UpdateMovie(db))
		movies.DELETE("/:id", ginMovieTrans.DeleteMovie(db))
	}
}