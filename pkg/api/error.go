package api

import (
	"fmt"
	"net/http"
)

func NewAPIError(statusCode int) APIError {
	return APIError{
		StatusCode: statusCode,
	}
}

type APIError struct {
	StatusCode int
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d %s", e.StatusCode, http.StatusText(e.StatusCode))
}

func NewUnauthorizedError() UnauthorizedError {
	return UnauthorizedError{
		APIError{
			StatusCode: http.StatusUnauthorized,
		},
	}
}

type UnauthorizedError struct {
	APIError
}
