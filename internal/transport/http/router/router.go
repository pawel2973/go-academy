package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/service/character"
	"github.com/pawel2973/go-academy/internal/service/movie"
	"github.com/pawel2973/go-academy/internal/transport/http/characters"
	"github.com/pawel2973/go-academy/internal/transport/http/movies"
)

// API aggregates HTTP handlers.
type API struct {
	MoviesController     *movies.MovieHandler
	CharactersController *characters.CharacterHandler
}

// NewAPI returns a new API.
func NewAPI(movieSvc *movie.MovieService, charSvc *character.CharacterService) *API {
	return &API{
		MoviesController:     movies.NewMovieHandler(movieSvc),
		CharactersController: characters.NewCharacterHandler(charSvc),
	}
}

// Register registers API routes.
func (a *API) Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	movies.RegisterRoutes(v1, a.MoviesController)
	characters.RegisterRoutes(v1, a.CharactersController)
}
