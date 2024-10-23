package middleware

import (
	"net/http"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequireAuth(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		jwtProvider := new(common.JWTProvider)

		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthoried)
			return
		}

		tokenString := strings.Split(bearerToken, " ")[1]

		token, err := jwtProvider.ValidateToken(tokenString)

		if err != nil || token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		// Store current use context
		if _, exists := c.Get("user"); !exists {
			// Extract claims from token
			claims, err := jwtProvider.ParseToken(tokenString)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, err)
				return
			}

			// Find and store user's current context
			storage := mssql.NewSQLStorage(db)
			business := business.NewGetUserBiz(storage)

			user, err := business.GetUserByEmail(c.Request.Context(), claims.Subject)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, err)
				return
			}

			c.Set("user", user)

			if model.IsAdmin(user.RoleCode) {
				c.Set("isAdmin", true)
			} else {
				c.Set("isAdmin", false)
			}
		}

		c.Next()
	}
}
