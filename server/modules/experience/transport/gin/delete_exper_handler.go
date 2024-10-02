package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteExperience(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		store := mssql.NewSQLStorage(db)
		biz := business.NewDeleteExperienceBiz(store)

		if err := biz.DeleteExperienceById(c, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
