package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/booking/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteBooking(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		store := mssql.NewSQLStorage(db)
		biz := business.NewDeleteBookingBiz(store)

		if err := biz.DeleteBookingById(c, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
