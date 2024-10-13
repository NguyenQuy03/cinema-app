package routes

import (
	"github.com/NguyenQuy03/cinema-app/server/middleware"
	ginTicketTypeTrans "github.com/NguyenQuy03/cinema-app/server/modules/ticketType/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupTicketTypeRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	ticketTypes := v1.Group("ticket_types", middleware.RequireAuth(db))
	{
		ticketTypes.POST("", ginTicketTypeTrans.CreateTicketType(db))
		ticketTypes.GET("/:id", ginTicketTypeTrans.GetTicketType(db))
		ticketTypes.PATCH("/:id", ginTicketTypeTrans.UpdateTicketType(db))
		ticketTypes.DELETE("/:id", ginTicketTypeTrans.DeleteTicketType(db))
	}
}
