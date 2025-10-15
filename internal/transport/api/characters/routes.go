package characters

import "github.com/labstack/echo/v4"

// RegisterRoutes registers character endpoints.
func RegisterRoutes(v1 *echo.Group, h *CharacterHandler) {
	r := v1.Group("/characters")
	r.GET("", h.NotImplemented)
	r.POST("", h.NotImplemented)
	r.DELETE("/:id", h.NotImplemented)
}
