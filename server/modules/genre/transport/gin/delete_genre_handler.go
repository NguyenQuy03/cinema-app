package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/genre/storage/sqlsv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		store := sqlsv.NewSQLStorage(db)
		biz := business.NewDeleteGenreBiz(store)

		if err := biz.DeleteGenreById(c, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
