package userRoutes

import (
	ginBookingTrans "github.com/NguyenQuy03/cinema-app/server/modules/booking/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupBookingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	bookings := v1.Group("bookings")
	{
		bookings.POST("", ginBookingTrans.CreateBooking(db))
		bookings.GET("", ginBookingTrans.ListBooking(db))
		bookings.GET("/:id", ginBookingTrans.GetBooking(db))
		bookings.PATCH("/:id", ginBookingTrans.UpdateBooking(db))
		bookings.DELETE("/:id", ginBookingTrans.DeleteBooking(db))
	}
}
