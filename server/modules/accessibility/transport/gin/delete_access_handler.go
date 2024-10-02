package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/accessibility/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAccessibility(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		store := mssql.NewSQLStorage(db)
		biz := business.NewDeleteAccessBiz(store)

		if err := biz.DeleteAccessById(c, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
