package openapi

// ServerImplementation aggregates two API modules so they jointly
// implement the ServerInterface generated from OpenAPI.
type ServerImplementation struct {
	*MovieAPI
	*CharacterAPI
}

// NewServerImplementation creates a new ServerImplementation from the
// provided MovieAPI and CharacterAPI.
func NewServerImplementation(movie *MovieAPI, character *CharacterAPI) *ServerImplementation {
	return &ServerImplementation{
		MovieAPI:     movie,
		CharacterAPI: character,
	}
}

// Compile-time assertion that ServerImplementation implements ServerInterface.
var _ ServerInterface = (*ServerImplementation)(nil)
