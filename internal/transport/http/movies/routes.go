package movies

import "github.com/labstack/echo/v4"

// RegisterRoutes registers movie endpoints.
func RegisterRoutes(v1 *echo.Group, h *MovieHandler) {
	r := v1.Group("/movies")
	r.GET("", h.NotImplemented)
	r.POST("", h.NotImplemented)
	r.PUT("/:id", h.NotImplemented)
	r.DELETE("/:id", h.NotImplemented)
}
