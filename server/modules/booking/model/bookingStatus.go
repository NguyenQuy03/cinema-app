package model

type BookingStatus struct {
	StatusCode string
	StatusName string
}

func (BookingStatus) TableName() string {
	return "booking_status"
}
