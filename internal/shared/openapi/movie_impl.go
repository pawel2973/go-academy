package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/modules/movie/domain"
	movieSvc "github.com/pawel2973/go-academy/internal/modules/movie/service"
	"github.com/pawel2973/go-academy/internal/shared/api"
	appErr "github.com/pawel2973/go-academy/internal/shared/errors"
)

// MovieAPI handles movie-related HTTP endpoints.
type MovieAPI struct {
	svc *movieSvc.MovieService
}

// NewMovieAPI creates a new MovieAPI instance.
func NewMovieAPI(svc *movieSvc.MovieService) *MovieAPI {
	return &MovieAPI{svc: svc}
}

// ListMovies handles GET /movies.
func (a *MovieAPI) ListMovies(ctx echo.Context) error {
	movies := a.svc.List()
	return ctx.JSON(http.StatusOK, movies)
}

// CreateMovie handles POST /movies.
func (a *MovieAPI) CreateMovie(ctx echo.Context) error {
	var req MovieRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}
	if req.Title == "" || req.Year <= 0 {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrMovieInvalidData)
	}
	m := a.svc.Create(req.Title, req.Year)
	return ctx.JSON(http.StatusCreated, m)
}

// GetMovie handles GET /movies/{movie_id}.
func (a *MovieAPI) GetMovie(ctx echo.Context, movieId string) error {
	if movieId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrIDRequired)
	}
	m, ok := a.svc.Get(movieId)
	if !ok {
		return ctx.JSON(http.StatusNotFound, appErr.ErrMovieNotFound)
	}
	return ctx.JSON(http.StatusOK, m)
}

// UpdateMovie handles PUT /movies/{movie_id}.
func (a *MovieAPI) UpdateMovie(ctx echo.Context, movieId string) error {
	if movieId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrIDRequired)
	}

	var req MovieRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}
	if req.Title == "" || req.Year <= 0 {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrMovieInvalidData)
	}

	movie := domain.Movie{
		ID:    movieId,
		Title: req.Title,
		Year:  req.Year,
	}

	updated, err := a.svc.Update(movie)
	if err != nil {
		return api.MapError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, updated)
}

// DeleteMovie handles DELETE /movies/{movie_id}.
func (a *MovieAPI) DeleteMovie(ctx echo.Context, movieId string) error {
	if movieId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrIDRequired)
	}
	if err := a.svc.Delete(movieId); err != nil {
		return api.MapError(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
