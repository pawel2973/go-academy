package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/service"
)

// ErrorResponse is a JSON error payload.
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewError creates an ErrorResponse with the given message.
func NewError(msg string) ErrorResponse {
	return ErrorResponse{Error: msg}
}

// MapError translates service errors to HTTP responses.
// Returns a JSON error with appropriate status code.
func MapError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrMovieNotFound):
		return c.JSON(http.StatusNotFound, NewError(err.Error()))
	case errors.Is(err, service.ErrIDRequired), errors.Is(err, service.ErrMovieIDRequired):
		return c.JSON(http.StatusBadRequest, NewError(err.Error()))
	default:
		return c.JSON(http.StatusInternalServerError, NewError("internal server error"))
	}
}
