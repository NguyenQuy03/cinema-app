package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginMovieTrans "github.com/NguyenQuy03/cinema-app/server/modules/movie/transport/gin"
	ginUserTrans "github.com/NguyenQuy03/cinema-app/server/modules/user/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupV1Router(router *gin.Engine, db *gorm.DB) *gin.Engine {
	v1 := router.Group("v1")
	{
		movieRoute(v1, db)
		userRoute(v1, db)
	}

	return router
}

func movieRoute(v1 *gin.RouterGroup, db *gorm.DB) {
	movie := v1.Group("movies", middleware.RequireAuth(db))
	{
		movie.POST("", ginMovieTrans.CreateMovie(db))
		movie.GET("", ginMovieTrans.ListMovie(db))
		movie.GET("/:id", ginMovieTrans.GetMovie(db))
		movie.PATCH("/:id", ginMovieTrans.UpdateMovie(db))
		movie.DELETE("/:id", ginMovieTrans.DeleteMovie(db))
	}
}

func userRoute(v1 *gin.RouterGroup, db *gorm.DB) {
	user := v1.Group("auth")
	{
		user.POST("/register", ginUserTrans.RegisterUser(db))
		user.POST("/login", ginUserTrans.AuthenticateUser(db))
	}
}
