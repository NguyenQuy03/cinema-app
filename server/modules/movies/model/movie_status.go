package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type MovieStatus int

const (
	MovieActiveStatus MovieStatus = iota
	MovieInActiveStatus
)

var allMovieStatuses = [2]string{"active", "inactive"}

func (status *MovieStatus) String() string {
	return allMovieStatuses[*status]
}

// Loop movie statuses slice to get the string value
func parseStrToMovieStatus(s string) (MovieStatus, error) {
	for i := range allMovieStatuses {
		if allMovieStatuses[i] == s {
			return MovieStatus(i), nil
		}
	}

	return MovieStatus(0), errors.New("invalid status string")
}

func (status *MovieStatus) Scan(value interface{}) error {
	// Assert that the value is of type string
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan data from SQL: expected string, got %T", value)
	}

	// Parse the string to your MovieStatus enum
	v, err := parseStrToMovieStatus(strValue)
	if err != nil {
		return fmt.Errorf("failed to parse status from SQL: %v", err)
	}

	// Set the parsed value to the status
	*status = v
	return nil
}

func (status *MovieStatus) Value() (driver.Value, error) {
	if status == nil {
		return nil, nil
	}

	return status.String(), nil
}

func (status *MovieStatus) MarshalJSON() ([]byte, error) {
	if status == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}

func (status *MovieStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	value, err := parseStrToMovieStatus(str)

	if err != nil {
		return err
	}

	*status = value

	return nil
}
