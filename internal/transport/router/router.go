package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/modules/character/service"
	"github.com/pawel2973/go-academy/internal/modules/character/transport/http"
	service2 "github.com/pawel2973/go-academy/internal/modules/movie/service"
	http2 "github.com/pawel2973/go-academy/internal/modules/movie/transport/http"
)

// API aggregates HTTP handlers.
type API struct {
	MoviesController     *http2.MovieHandler
	CharactersController *http.CharacterHandler
}

// NewAPI returns a new API.
func NewAPI(movieSvc *service2.MovieService, charSvc *service.CharacterService) *API {
	return &API{
		MoviesController:     http2.NewMovieHandler(movieSvc),
		CharactersController: http.NewCharacterHandler(charSvc),
	}
}

// Register registers API routes.
func (a *API) Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	RegisterMovieRoutes(v1, a.MoviesController)
	RegisterCharacterRoutes(v1, a.CharactersController)
}
