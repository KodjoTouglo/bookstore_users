package errors

import "net/http"

type APIError struct {
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	StatusCode int    `json:"status_code"`
}

func BadRequestError(detail string, message ...string) *APIError {
	apiErr := &APIError{
		Detail:     detail,
		StatusCode: http.StatusBadRequest,
	}
	if len(message) > 0 {
		apiErr.Message = message[0]
	}
	return apiErr
}

func NotFoundError(detail string, message ...string) *APIError {
	apiErr := &APIError{
		Detail:     detail,
		StatusCode: http.StatusNotFound,
	}
	if len(message) > 0 {
		apiErr.Message = message[0]
	}
	return apiErr
}

func InternalServerError(detail string, message ...string) *APIError {
	apiErr := &APIError{
		Detail:     detail,
		StatusCode: http.StatusInternalServerError,
	}
	if len(message) > 0 {
		apiErr.Message = message[0]
	}
	return apiErr
}
