package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type UserRole int

const (
	Customer UserRole = iota
	Cashier
	Manager
)

var allUserRoles = [3]string{"customer", "cashier", "manager"}

func (role *UserRole) String() string {
	return allUserRoles[*role]
}

// Loop movie statuses slice to get the string value
func parseStrToUserRole(s string) (UserRole, error) {
	for i := range allUserRoles {
		if allUserRoles[i] == s {
			return UserRole(i), nil
		}
	}

	return UserRole(0), errors.New("invalid role string")
}

func (role *UserRole) Scan(value interface{}) error {
	// Assert that the value is of type string
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan data from SQL: expected string, got %T", value)
	}

	// Parse the string to your user role enum
	v, err := parseStrToUserRole(strValue)
	if err != nil {
		return fmt.Errorf("failed to parse role from SQL: %v", err)
	}

	// Set the parsed value to the role
	*role = v
	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}

	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	if role == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

func (role *UserRole) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	value, err := parseStrToUserRole(str)

	if err != nil {
		return err
	}

	*role = value

	return nil
}
