package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawel2973/go-academy/internal/modules/character/service"
)

// CharacterHandler is a controller responsible for the Character resource.
type CharacterHandler struct {
	svc *service.CharacterService
}

// NewCharacterHandler creates a new CharacterHandler.
func NewCharacterHandler(svc *service.CharacterService) *CharacterHandler {
	return &CharacterHandler{svc: svc}
}

// NotImplemented handles unimplemented endpoints.
func (h *CharacterHandler) NotImplemented(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
