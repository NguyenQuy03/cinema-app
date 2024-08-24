package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrResponse(statusCode int, rootErr error, message, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(rootErr error, message, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(rootErr error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, message string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, message, root.Error(), key)
	}

	return NewErrorResponse(errors.New(message), message, message, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrResponse(http.StatusInternalServerError, err, "something went wrong from DB", err.Error(), "DB_ERROR")
}

func ErrInvalidReq(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQUEST_ERROR")
}

func ErrInternal(err error) *AppError {
	return NewFullErrResponse(http.StatusInternalServerError, err, "some thing went wrong in the server", err.Error(), "INTERNAL_ERROR")
}

func ErrCannotListEntity(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot list %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_LIST_%s_ERROR", strings.ToUpper(entity)))
}

func ErrCannotDeleteEntity(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_DELETE_%s_ERROR", strings.ToUpper(entity)))
}

func ErrCannotUpdateEntity(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot update %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_UPDATE_%s_ERROR", strings.ToUpper(entity)))
}

func ErrCannotCreateEntity(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot create %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_CREATE_%s_ERROR", strings.ToUpper(entity)))
}

func ErrCannotGetEntity(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot get %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_GET_%s_ERROR", strings.ToUpper(entity)))
}

func ErrEntityDeleted(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s is deleted", strings.ToLower(entity)), fmt.Sprintf("%s_DELETED_ERROR", strings.ToUpper(entity)))
}

func ErrEntityNotFound(err error, entity string) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s not found", strings.ToLower(entity)), fmt.Sprintf("%s_NOT_FOUND_ERROR", strings.ToUpper(entity)))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, "you have no permission", "NO_PERMISSION_ERROR")
}

// Using recursive to get the root error
// If current error does not have root error -> call RootError() func of inside error
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

var ErrRecordNotFound = errors.New("record not found")
