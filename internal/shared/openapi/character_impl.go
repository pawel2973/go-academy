package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	charSvc "github.com/pawel2973/go-academy/internal/modules/character/service"
	"github.com/pawel2973/go-academy/internal/shared/api"
	appErr "github.com/pawel2973/go-academy/internal/shared/errors"
)

// CharacterAPI handles character-related HTTP endpoints.
type CharacterAPI struct {
	svc *charSvc.CharacterService
}

// NewCharacterAPI creates a new CharacterAPI instance.
func NewCharacterAPI(svc *charSvc.CharacterService) *CharacterAPI {
	return &CharacterAPI{svc: svc}
}

// ListCharacters handles GET /characters.
func (a *CharacterAPI) ListCharacters(ctx echo.Context) error {
	chars := a.svc.ListAll()
	return ctx.JSON(http.StatusOK, chars)
}

// ListCharactersForMovie handles GET /movies/{movie_id}/characters.
func (a *CharacterAPI) ListCharactersForMovie(ctx echo.Context, movieId string) error {
	if movieId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrMovieIDRequired)
	}
	chars := a.svc.ListByMovie(movieId)
	return ctx.JSON(http.StatusOK, chars)
}

// CreateCharacterForMovie handles POST /movies/{movie_id}/characters.
func (a *CharacterAPI) CreateCharacterForMovie(ctx echo.Context, movieId string) error {
	if movieId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrMovieIDRequired)
	}

	var req CharacterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}
	if req.Name == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}

	ch, err := a.svc.Create(movieId, req.Name)
	if err != nil {
		return api.MapError(ctx, err)
	}
	return ctx.JSON(http.StatusCreated, ch)
}

// GetCharacter handles GET /characters/{character_id}.
func (a *CharacterAPI) GetCharacter(ctx echo.Context, characterId string) error {
	if characterId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrIDRequired)
	}
	ch, ok := a.svc.Get(characterId)
	if !ok {
		return ctx.JSON(http.StatusNotFound, appErr.ErrCharacterNotFound)
	}
	return ctx.JSON(http.StatusOK, ch)
}

// UpdateCharacter handles PUT /characters/{character_id}.
func (a *CharacterAPI) UpdateCharacter(ctx echo.Context, characterId string) error {
	if characterId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrIDRequired)
	}

	var req CharacterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}
	if req.Name == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrInvalidJSON)
	}

	newMovieID := ""
	if req.MovieId != nil {
		newMovieID = *req.MovieId
	}

	ch, err := a.svc.UpdateDomain(characterId, newMovieID, req.Name)
	if err != nil {
		return api.MapError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, ch)
}

// DeleteCharacter handles DELETE /characters/{character_id}.
func (a *CharacterAPI) DeleteCharacter(ctx echo.Context, characterId string) error {
	if characterId == "" {
		return ctx.JSON(http.StatusBadRequest, appErr.ErrIDRequired)
	}
	if err := a.svc.Delete(characterId); err != nil {
		return api.MapError(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
