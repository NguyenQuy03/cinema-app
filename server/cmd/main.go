package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MovieStatus int

const (
	MovieActiveStatus MovieStatus = iota
	MovieInActiveStatus
)

var allMovieStatuses = [2]string{"active", "inactive"}

func (movieStatus MovieStatus) String() string {
	return allMovieStatuses[movieStatus]
}

func parseStrToMovieStatus(s string) (MovieStatus, error) {
	for i := range allMovieStatuses {
		if allMovieStatuses[i] == s {
			return MovieStatus(i), nil
		}
	}

	return MovieStatus(0), errors.New("invalid status string")
}

func (status *MovieStatus) Scan(value interface{}) error {
	// Assert that the value is of type string
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan data from SQL: expected string, got %T", value)
	}

	// Parse the string to your MovieStatus enum
	v, err := parseStrToMovieStatus(strValue)
	if err != nil {
		return fmt.Errorf("failed to parse status from SQL: %v", err)
	}

	// Set the parsed value to the status
	*status = v
	return nil
}

func (status *MovieStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}

type Movie struct {
	Id          int          `json:"-" gorm:"column:id"`
	Description string       `json:"description" gorm:"column:description"`
	Duration    int          `json:"duration" gorm:"column:duration"`
	Genre       string       `json:"genre" gorm:"column:genre"`
	TrailerLink string       `json:"trailer_link" gorm:"column:trailer_link"`
	Status      *MovieStatus `json:"status" gorm:"column:status"`
}

func (Movie) TableName() string { return "movies" }

type MovieCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Description string `json:"description" gorm:"column:description"`
	Duration    int    `json:"duration" gorm:"column:duration"`
	Genre       string `json:"genre" gorm:"column:genre"`
	TrailerLink string `json:"trailer_link" gorm:"column:trailer_link"`
}

func (MovieCreation) TableName() string { return "movies" }

type MovieUpdate struct {
	Description *string `json:"description" gorm:"column:description"`
	Duration    int     `json:"duration" gorm:"column:duration"`
	Genre       *string `json:"genre" gorm:"column:genre"`
	TrailerLink *string `json:"trailer_link" gorm:"column:trailer_link"`
}

func (MovieUpdate) TableName() string { return "movies" }

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Page = 5
	}
}

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
			movies.POST("", CreateMovie(db))
			movies.GET("", ListMovie(db))
			movies.GET("/:id", GetMovie(db))
			movies.PATCH("/:id", UpdateMovie(db))
			movies.DELETE("/:id", DeleteMovie(db))
		}
	}

	// Define a route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")

}

func CreateMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data MovieCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}

func GetMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data Movie

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func UpdateMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data MovieUpdate

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Table(Movie{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "inactive",
		}).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func ListMovie(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var p Paging

		if err := ctx.ShouldBind(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		p.Process()

		var result []Movie

		if err := db.Table(Movie{}.TableName()).Count(&p.Total).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Order("id desc").Offset((p.Page - 1) * p.Limit).Limit(p.Limit).Find(&result).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": p,
		})
	}
}
