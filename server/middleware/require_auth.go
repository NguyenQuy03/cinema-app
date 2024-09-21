package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RequireAuth(db *gorm.DB, redisDB *redis.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			err := common.NewUnauthorized(errors.New("missing or invalid token"), "Authorization header is missing or invalid", "TOKEN_MISSING_OR_INVALID_ERR")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		tokenString := strings.Split(bearerToken, " ")[1]

		handlerTokenBiz := business.HandleTokenBiz{}
		token, err := handlerTokenBiz.ValidateToken(tokenString)

		if err != nil || token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		// Store current use context
		if _, exists := c.Get("user"); !exists {
			// Extract email from token
			email, err := handlerTokenBiz.ExtractEmail(token)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, err)
				return
			}

			// Find and store user's current context
			storage := sqlsv.NewSQLStorage(db)
			business := business.NewGetUserBiz(storage)

			user, err := business.GetUserByEmail(c.Request.Context(), email)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, err)
				return
			}

			c.Set("user", user)
		}

		c.Next()
	}
}
