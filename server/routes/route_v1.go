package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginUserTrans "github.com/NguyenQuy03/cinema-app/server/modules/auth/transport/gin"
	ginGenreTrans "github.com/NguyenQuy03/cinema-app/server/modules/genre/transport/gin"
	ginMovieTrans "github.com/NguyenQuy03/cinema-app/server/modules/movie/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupV1Router(router *gin.Engine, db *gorm.DB, redisDB *redis.Client) *gin.Engine {
	v1 := router.Group("v1")
	{
		setupAuthRoutes(v1, db, redisDB)

		setupMovieRoutes(v1, db, redisDB)
		setupGenreRoutes(v1, db, redisDB)
	}

	return router
}

func setupAuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	auth := v1.Group("auth")
	{
		auth.POST("/register", ginUserTrans.RegisterUser(db))
		auth.POST("/login", ginUserTrans.AuthenticateUser(db, redisDB))
		auth.POST("/refresh-token", ginUserTrans.RefreshToken(db, redisDB))
	}
}

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

func setupGenreRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	genres := v1.Group("genres", middleware.RequireAuth(db, redisDB))
	{
		genres.POST("", ginGenreTrans.CreateGenre(db))
		// genres.GET("", ginMovieTrans.ListMovie(db))
		genres.GET("/:id", ginGenreTrans.GetGenre(db))
		genres.PATCH("/:id", ginGenreTrans.UpdateGenre(db))
		genres.DELETE("/:id", ginGenreTrans.DeleteGenre(db))
	}
}
