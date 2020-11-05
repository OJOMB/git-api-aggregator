package utils

import "net/http"

// APIError defines the method set of behavious for a valid API error
type APIError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (ae *apiError) GetError() string {
	return ae.Error
}

func (ae *apiError) GetStatus() int {
	return ae.Status
}

func (ae *apiError) GetMessage() string {
	return ae.Message
}

// NewAPIError returns a new APIError
func NewAPIError(status int, message string) APIError {
	return &apiError{
		Status:  status,
		Message: message,
	}
}

// NewNotFoundError returns a 404 type API error
func NewNotFoundError(msg string) APIError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

// NewInternalServiceError returns a 500 type API error
func NewInternalServiceError(msg string) APIError {
	return &apiError{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// NewBadRequestError returns a 400 type API error
func NewBadRequestError(msg string) APIError {
	return &apiError{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}
