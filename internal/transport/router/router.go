package router

import (
	"github.com/labstack/echo/v4"
	characterSvc "github.com/pawel2973/go-academy/internal/modules/character/service"
	characterHTTP "github.com/pawel2973/go-academy/internal/modules/character/transport/http"
	movieSvc "github.com/pawel2973/go-academy/internal/modules/movie/service"
	movieHTTP "github.com/pawel2973/go-academy/internal/modules/movie/transport/http"
)

// API aggregates HTTP handlers.
type API struct {
	MoviesController     *movieHTTP.MovieHandler
	CharactersController *characterHTTP.CharacterHandler
}

// NewAPI returns a new API.
func NewAPI(movieSvc *movieSvc.MovieService, charSvc *characterSvc.CharacterService) *API {
	return &API{
		MoviesController:     movieHTTP.NewMovieHandler(movieSvc),
		CharactersController: characterHTTP.NewCharacterHandler(charSvc),
	}
}

// Register registers API routes.
func (a *API) Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	RegisterMovieRoutes(v1, a.MoviesController)
	RegisterCharacterRoutes(v1, a.CharactersController)
}
