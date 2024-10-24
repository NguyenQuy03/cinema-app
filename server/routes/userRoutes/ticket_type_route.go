package userRoutes

import (
	ginTicketTypeTrans "github.com/NguyenQuy03/cinema-app/server/modules/ticketType/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupTicketTypeRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	ticketTypes := v1.Group("ticket-types")
	{
		ticketTypes.GET("", ginTicketTypeTrans.ListTicketType(db))
	}
}
