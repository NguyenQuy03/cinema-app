package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/sqlsv"
	"github.com/NguyenQuy03/cinema-app/server/utils/jwtUtil"
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

		token, err := jwtUtil.ValidateToken(tokenString)

		if err != nil || token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		// Extract email from token
		email, err := jwtUtil.ExtractEmail(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		// Compare with existed access_token is redis
		key := fmt.Sprintf("user_sessions:%s", email)
		res, err := redisDB.HGetAll(c.Request.Context(), key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		if !strings.EqualFold(res["access_token"], tokenString) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrInvalidToken)
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

		if _, exists := c.Get("user"); !exists {
			c.Set("user", user)
		}

		c.Next()
	}
}
