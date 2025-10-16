package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	appErr "github.com/pawel2973/go-academy/internal/shared/errors"
)

// ErrorResponse is a JSON error payload.
type ErrorResponse struct {
	Error string `json:"error"`
}

// newError creates an ErrorResponse with the given message.
func newError(msg string) ErrorResponse {
	return ErrorResponse{Error: msg}
}

// MapError translates service errors to HTTP responses.
// Returns a JSON error with the appropriate status code.
func MapError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, appErr.ErrMovieNotFound),
		errors.Is(err, appErr.ErrCharacterNotFound):
		return c.JSON(http.StatusNotFound, newError(err.Error()))
	case errors.Is(err, appErr.ErrInvalidJSON),
		errors.Is(err, appErr.ErrMovieInvalidData),
		errors.Is(err, appErr.ErrIDRequired),
		errors.Is(err, appErr.ErrMovieIDRequired):
		return c.JSON(http.StatusBadRequest, newError(err.Error()))
	default:
		return c.JSON(http.StatusInternalServerError, newError("internal server error"))
	}
}
