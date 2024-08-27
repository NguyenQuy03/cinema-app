package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/user/storage/sqlsv"
	"github.com/NguyenQuy03/cinema-app/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RequireAuth(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			err := common.NewUnauthorized(errors.New("missing or invalid token"), "Authorization header is missing or invalid", "Authorization header is missing or invalid", "TOKEN_MISSING_OR_INVALID_ERR")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		tokenString := strings.Split(bearerToken, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check algorithm used for generating token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.Abort()
				return nil, common.NewUnauthorized(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]), "unexpected signing method", "unexpected signing method", "SIGNING_METHOD_ERR")
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			var returnErr error

			switch {
			case errors.Is(err, jwt.ErrSignatureInvalid):
				returnErr = model.ErrInvalidToken
			case errors.Is(err, jwt.ErrTokenMalformed):
				returnErr = model.ErrMalformedToken
			case errors.Is(err, jwt.ErrTokenExpired):
				returnErr = model.ErrExpirededToken
			default:
				// Generic error handling for other cases
				returnErr = model.ErrInvalidToken
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, returnErr)
			return
		}

		// Extract email from token
		email, err := utils.ExtractEmail(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		storage := sqlsv.NewSQLStorage(db)
		business := business.NewGetUserBiz(storage)

		user, err := business.GetUserByEmail(c.Request.Context(), email)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
