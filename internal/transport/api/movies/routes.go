package movies

import "github.com/labstack/echo/v4"

// RegisterRoutes registers movie endpoints.
func RegisterRoutes(v1 *echo.Group, h *MovieHandler) {
	r := v1.Group("/movies")
	r.GET("", h.ListMovies)
	r.GET("/:id", h.GetMovie)
	r.POST("", h.CreateMovie)
	r.PUT("/:id", h.UpdateMovie)
	r.DELETE("/:id", h.DeleteMovie)
}
