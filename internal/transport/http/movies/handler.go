package movies

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/service/movie"
)

// MovieHandler is a controller responsible for the Movie resource.
type MovieHandler struct {
	svc *movie.MovieService
}

// NewMovieHandler creates a new MovieHandler.
func NewMovieHandler(svc *movie.MovieService) *MovieHandler {
	return &MovieHandler{svc: svc}
}

// NotImplemented handles unimplemented endpoints.
func (h *MovieHandler) NotImplemented(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
