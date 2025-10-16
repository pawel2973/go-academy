package router

import (
	"github.com/labstack/echo/v4"
	movieHTTP "github.com/pawel2973/go-academy/internal/modules/movie/transport/http"
)

// RegisterMovieRoutes registers movie endpoints.
func RegisterMovieRoutes(v1 *echo.Group, h *movieHTTP.MovieHandler) {
	r := v1.Group("/movies")
	r.GET("", h.ListMovies)
	r.GET("/:id", h.GetMovie)
	r.POST("", h.CreateMovie)
	r.PUT("/:id", h.UpdateMovie)
	r.DELETE("/:id", h.DeleteMovie)
}
