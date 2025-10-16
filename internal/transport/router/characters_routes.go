package router

import (
	"github.com/labstack/echo/v4"
	characterHTTP "github.com/pawel2973/go-academy/internal/modules/character/transport/http"
)

// RegisterCharacterRoutes registers character endpoints.
func RegisterCharacterRoutes(v1 *echo.Group, h *characterHTTP.CharacterHandler) {
	r := v1.Group("/characters")
	r.GET("", h.NotImplemented)
	r.POST("", h.NotImplemented)
	r.DELETE("/:id", h.NotImplemented)
}
