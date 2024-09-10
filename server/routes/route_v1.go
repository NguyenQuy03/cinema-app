package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginUserTrans "github.com/NguyenQuy03/cinema-app/server/modules/auth/transport/gin"
	ginMovieTrans "github.com/NguyenQuy03/cinema-app/server/modules/movie/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupV1Router(router *gin.Engine, db *gorm.DB, redisDB *redis.Client) *gin.Engine {
	v1 := router.Group("v1")
	{
		setupMovieRoutes(v1, db, redisDB)
		setupAuthRoutes(v1, db, redisDB)
	}

	return router
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

func setupAuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redisDB *redis.Client) {
	auth := v1.Group("auth")
	{
		auth.POST("/register", ginUserTrans.RegisterUser(db))
		auth.POST("/login", ginUserTrans.AuthenticateUser(db, redisDB))
		auth.POST("/refresh-token", ginUserTrans.RefreshToken(db, redisDB))
	}
}
