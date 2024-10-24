package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdminRole() func(*gin.Context) {
	return func(c *gin.Context) {
		isAdmin, isExist := c.Get("isAdmin")
		if !isExist || isAdmin == false {
			c.AbortWithStatusJSON(http.StatusForbidden, ErrNoPermission)
			return
		}

		c.Next()
	}
}
