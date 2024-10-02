package gintrans

import (
	"net/http"
	"strconv"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateExperience(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.ExperienceUpdate
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewUpdateExperienceBiz(storage)

		if err := business.UpdateExperience(ctx, id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(true))
	}
}
