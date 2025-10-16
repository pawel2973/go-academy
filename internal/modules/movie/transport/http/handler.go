package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/modules/movie/domain"
	"github.com/pawel2973/go-academy/internal/modules/movie/service"
	"github.com/pawel2973/go-academy/internal/shared/api"
	appErr "github.com/pawel2973/go-academy/internal/shared/errors"
)

// MovieHandler is a controller responsible for the Movie resource.
type MovieHandler struct {
	svc *service.MovieService
}

// NewMovieHandler creates a new MovieHandler.
func NewMovieHandler(svc *service.MovieService) *MovieHandler {
	return &MovieHandler{svc: svc}
}

// MovieRequest represents the JSON payload used to create or update a movie.
type MovieRequest struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

// MovieResponse represents the JSON response returned to the client.
type MovieResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

// ListMovies handles GET /api/v1/movies.
// It returns all movies as a JSON array.
func (h *MovieHandler) ListMovies(c echo.Context) error {
	movies := h.svc.List()
	resp := make([]MovieResponse, 0, len(movies))
	for _, m := range movies {
		resp = append(resp, MovieResponse{
			ID:    m.ID,
			Title: m.Title,
			Year:  m.Year,
		})
	}
	return c.JSON(http.StatusOK, resp)
}

// GetMovie handles GET /api/v1/movies/:id.
// It returns a single movie by its ID.
func (h *MovieHandler) GetMovie(c echo.Context) error {
	id := c.Param("id")
	m, ok := h.svc.Get(id)
	if !ok {
		return c.JSON(http.StatusNotFound, appErr.ErrMovieNotFound)
	}
	resp := MovieResponse{ID: m.ID, Title: m.Title, Year: m.Year}
	return c.JSON(http.StatusOK, resp)
}

// CreateMovie handles POST /api/v1/movies.
// It creates a new movie using data from the request body.
func (h *MovieHandler) CreateMovie(c echo.Context) error {
	var req MovieRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}
	if req.Title == "" || req.Year <= 0 {
		return c.JSON(http.StatusBadRequest, appErr.ErrMovieInvalidData)
	}

	m := h.svc.Create(req.Title, req.Year)
	resp := MovieResponse{ID: m.ID, Title: m.Title, Year: m.Year}
	return c.JSON(http.StatusCreated, resp)
}

// UpdateMovie handles PUT /api/v1/movies/:id.
// It updates an existing movie using the provided ID and request body.
func (h *MovieHandler) UpdateMovie(c echo.Context) error {
	id := c.Param("id")
	var req MovieRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}

	m := domain.Movie{ID: id, Title: req.Title, Year: req.Year}
	updated, err := h.svc.Update(m)
	if err != nil {
		return api.MapError(c, err)
	}

	resp := MovieResponse{ID: updated.ID, Title: updated.Title, Year: updated.Year}
	return c.JSON(http.StatusOK, resp)
}

// DeleteMovie handles DELETE /api/v1/movies/:id.
// It deletes the specified movie and returns 204 No Content on success.
func (h *MovieHandler) DeleteMovie(c echo.Context) error {
	id := c.Param("id")
	if err := h.svc.Delete(id); err != nil {
		return api.MapError(c, err)
	}
	return c.NoContent(http.StatusNoContent)
}
