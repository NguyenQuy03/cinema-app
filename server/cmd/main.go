package main

import (
	"log"
	"os"

	"github.com/NguyenQuy03/cinema-app/server/configs"
	"github.com/NguyenQuy03/cinema-app/server/db"
	"github.com/NguyenQuy03/cinema-app/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	// Initialize the database connection
	db, err := db.InitSQLServerDB()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Setup the router
	router := gin.Default()

	// Setup v1 routes
	routes.SetupV1Router(router, db)

	// Run the server on port 8080
	router.Run(":" + os.Getenv("PORT"))
}
