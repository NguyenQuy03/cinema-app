package business

import (
	"context"

	"github.com/NguyenQuy03/cinema-app/server/common"
	bookingModel "github.com/NguyenQuy03/cinema-app/server/modules/booking/model"
	seatModel "github.com/NguyenQuy03/cinema-app/server/modules/seat/model"
	"gorm.io/gorm"
)

type CreateBookingStorage interface {
	CreateBooking(ctx context.Context, data *bookingModel.BookingCreation) error
	CreateBookingSeat(ctx context.Context, data []*bookingModel.BookingSeatCreation) error
	CreateBookingTicket(ctx context.Context, data []*bookingModel.BookingTicketCreation) error

	UpdateSeatInBulk(ctx context.Context, seatIds []int, field string, value string) error

	Begin() *gorm.DB
}

type SlugProvider interface {
	GenerateSlug(input string) string
}

type createBookingBiz struct {
	storage CreateBookingStorage
}

func NewCreateBookingBiz(storage CreateBookingStorage) *createBookingBiz {
	return &createBookingBiz{
		storage: storage,
	}
}

func (biz *createBookingBiz) CreateBooking(ctx context.Context, data *bookingModel.BookingCreation) error {
	// Validate input data at the start
	if data == nil || len(data.Seats) == 0 || len(data.Tickets) == 0 {
		return bookingModel.ErrInvalidInput
	}

	// Start a new transaction
	tx := biz.storage.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			// If no panic, check if the transaction should be committed or already rolled back
			if tx != nil {
				tx.Rollback()
			}
		}
	}()
	// Create the booking
	if err := tx.Table("booking").Create(data).Error; err != nil {
		tx.Rollback()
		return common.ErrCannotCreateEntity(err, bookingModel.BookingEntityName)
	}

	// Handle Booking_Seat
	bookingSeats := make([]*bookingModel.BookingSeatCreation, 0, len(data.Seats))
	for _, seatId := range data.Seats {
		bookingSeats = append(bookingSeats, &bookingModel.BookingSeatCreation{
			BookingId: data.Id,
			SeatId:    seatId,
		})
	}

	if err := tx.Create(bookingSeats).Error; err != nil {
		tx.Rollback()
		return common.ErrCannotCreateEntity(err, bookingModel.BookingEntityName)
	}

	// Update seat status to reserved
	reservedStatus := seatModel.SeatReservedStatus
	if err := biz.storage.UpdateSeatInBulk(ctx, data.Seats, "status", reservedStatus.String()); err != nil {
		tx.Rollback()
		return common.ErrCannotCreateEntity(err, bookingModel.BookingEntityName)
	}

	// Handle Booking_Ticket
	bookingTickets := make([]*bookingModel.BookingTicketCreation, 0, len(data.Tickets))
	for _, ticket := range data.Tickets {
		// Validate ticket fields
		if ticket.TicketTypeId <= 0 || ticket.TicketQuantity <= 0 {
			tx.Rollback()
			return bookingModel.ErrInvalidInput
		}

		bookingTickets = append(bookingTickets, &bookingModel.BookingTicketCreation{
			BookingId:      data.Id,
			TicketTypeId:   ticket.TicketTypeId,
			TicketQuantity: ticket.TicketQuantity,
		})
	}

	if err := tx.Create(bookingTickets).Error; err != nil {
		tx.Rollback()
		return common.ErrCannotCreateEntity(err, bookingModel.BookingEntityName)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
