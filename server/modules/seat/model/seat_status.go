package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type SeatStatus int

const (
	SeatReservedStatus SeatStatus = iota
	SeatAvailableStatus
	SeatDisabledStatus
)

var allSeatStatuses = [3]string{"reserved", "available", "disabled"}

func (status *SeatStatus) String() string {
	return allSeatStatuses[*status]
}

// Loop Seat statuses slice to get the string value
func parseStrToSeatStatus(s string) (SeatStatus, error) {
	for i := range allSeatStatuses {
		if allSeatStatuses[i] == s {
			return SeatStatus(i), nil
		}
	}

	return SeatStatus(0), errors.New("invalid status string")
}

func (status *SeatStatus) Scan(value interface{}) error {
	// Assert that the value is of type string
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan data from SQL: expected string, got %T", value)
	}

	// Parse the string to your SeatStatus enum
	v, err := parseStrToSeatStatus(strValue)
	if err != nil {
		return fmt.Errorf("failed to parse status from SQL: %v", err)
	}

	// Set the parsed value to the status
	*status = v
	return nil
}

func (status *SeatStatus) Value() (driver.Value, error) {
	if status == nil {
		return nil, nil
	}

	return status.String(), nil
}

func (status *SeatStatus) MarshalJSON() ([]byte, error) {
	if status == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}

func (status *SeatStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	value, err := parseStrToSeatStatus(str)

	if err != nil {
		return err
	}

	*status = value

	return nil
}
