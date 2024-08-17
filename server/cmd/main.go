package main

import (
	"fmt"
	"os"
	"path/filepath"

	gintrans "github.com/NguyenQuy03/cinema-app/server/modules/movies/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	// Connect to DB
	err := godotenv.Load(filepath.Join("../", ".env"))

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	router := gin.Default()

	v1 := router.Group("v1")
	{
		movies := v1.Group("movies")
		{
			movies.POST("", gintrans.CreateMovie(db))
			movies.GET("", gintrans.ListMovie(db))
			movies.GET("/:id", gintrans.GetMovie(db))
			movies.PATCH("/:id", gintrans.UpdateMovie(db))
			movies.DELETE("/:id", gintrans.DeleteMovie(db))
		}
	}

	// Run the server on port 8080
	router.Run(":8080")

}
