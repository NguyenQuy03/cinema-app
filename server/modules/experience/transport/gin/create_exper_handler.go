package gintrans

import (
	"net/http"

	"github.com/NguyenQuy03/cinema-app/server/common"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/business"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/model"
	"github.com/NguyenQuy03/cinema-app/server/modules/experience/storage/mssql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExperience(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.ExperienceCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		storage := mssql.NewSQLStorage(db)
		business := business.NewCreateExperienceBiz(storage)

		if err := business.CreateExperience(ctx, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidReq(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSimpleAppResponse(data.Id))
	}
}
