package errs

import "net/http"

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message" example:"Something went wrong"`
	Log     string `json:"-"`
}

func (err *AppError) Error() string {
	return err.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message, log string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Log:     log,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewAuthenthorizationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewTooManyRequestsError(message string) *AppError {
	return &AppError{
		Code:    http.StatusTooManyRequests,
		Message: message,
	}
}
