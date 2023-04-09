package errors

import "net/http"

type APIError struct {
	Detail     string `json:"detail"`
	StatusCode int    `json:"status_code"`
}

func BadRequestError(detail string) *APIError {
	return &APIError{
		Detail:     detail,
		StatusCode: http.StatusBadRequest,
	}
}

func NotFoundError(detail string) *APIError {
	return &APIError{
		Detail:     detail,
		StatusCode: http.StatusNotFound,
	}
}

func InternalServerError(detail string) *APIError {
	return &APIError{
		Detail:     detail,
		StatusCode: http.StatusInternalServerError,
	}
}
