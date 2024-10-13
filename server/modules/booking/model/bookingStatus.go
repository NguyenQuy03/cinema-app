package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type BookingStatus int

const (
	BookingPendingStatus BookingStatus = iota
	BookingConfirmedStatus
	BookingCanceledStatus
)

var allBookingStatuses = [3]string{"pending", "confirmed", "canceled"}

func (status *BookingStatus) String() string {
	return allBookingStatuses[*status]
}

// Loop movie statuses slice to get the string value
func parseStrToBookingStatus(s string) (BookingStatus, error) {
	for i := range allBookingStatuses {
		if allBookingStatuses[i] == s {
			return BookingStatus(i), nil
		}
	}

	return BookingStatus(0), errors.New("invalid status string")
}

func (status *BookingStatus) Scan(value interface{}) error {
	// Assert that the value is of type string
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan data from SQL: expected string, got %T", value)
	}

	// Parse the string to your BookingStatus enum
	v, err := parseStrToBookingStatus(strValue)
	if err != nil {
		return fmt.Errorf("failed to parse status from SQL: %v", err)
	}

	// Set the parsed value to the status
	*status = v
	return nil
}

func (status *BookingStatus) Value() (driver.Value, error) {
	if status == nil {
		return nil, nil
	}

	return status.String(), nil
}

func (status *BookingStatus) MarshalJSON() ([]byte, error) {
	if status == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}

func (status *BookingStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	value, err := parseStrToBookingStatus(str)

	if err != nil {
		return err
	}

	*status = value

	return nil
}
