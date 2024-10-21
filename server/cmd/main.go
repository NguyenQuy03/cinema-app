package main

import (
	"log"
	"os"

	"github.com/NguyenQuy03/cinema-app/server/configs"
	"github.com/NguyenQuy03/cinema-app/server/db"
	"github.com/NguyenQuy03/cinema-app/server/routes/adminRoutes"
	"github.com/NguyenQuy03/cinema-app/server/routes/commonRoutes"
	"github.com/NguyenQuy03/cinema-app/server/routes/userRoutes"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	// Initialize the application
	sqlserverDB, redisDB, cld := initializeServices()

	// Setup the router
	router := gin.Default()
	setupRoutes(router, sqlserverDB, redisDB, cld)

	// Run the server on port 8080
	router.Run(":" + os.Getenv("PORT"))
}

func initializeServices() (*gorm.DB, *redis.Client, *cloudinary.Cloudinary) {
	// Initialize the database connection
	sqlserverDB, err := db.InitSQLServerDB()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Init Redis DB
	opt, err := db.InitRedisDB()
	if err != nil {
		panic(err)
	}
	redisDB := redis.NewClient(opt)

	// Init Cloudinary Storage
	cld, err := db.InitCloudinaryStorage()
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	return sqlserverDB, redisDB, cld
}

func setupRoutes(router *gin.Engine, sqlserverDB *gorm.DB, redisDB *redis.Client, cld *cloudinary.Cloudinary) {
	// Setup v1 common routes
	commonRoutes.SetupCommonV1Router(router, sqlserverDB, redisDB, cld)

	// Setup v1 user routes
	userRoutes.SetupV1Router(router, sqlserverDB, redisDB)

	// Setup v1 admin routes
	adminRoutes.SetupAdminV1Router(router, sqlserverDB)
}
