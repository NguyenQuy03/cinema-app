package gin

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/auth/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.UserRegister

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		email := strings.TrimSpace(data.Email)
		password := data.Password

		// Validate email
		_, err := mail.ParseAddress(email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrEmailInvalid)
			return
		}

		// Validate password
		if len(password) < 6 {
			ctx.JSON(http.StatusBadRequest, model.ErrShortPass)
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewRegisterUserBiz(storage)

		// Register user
		if err := business.RegisterUser(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
