package errors

import "errors"

var (
	// General errors
	ErrIDRequired  = errors.New("id required")
	ErrInvalidJSON = errors.New("invalid JSON")

	// Movie domain errors
	ErrMovieNotFound    = errors.New("movie not found")
	ErrMovieInvalidData = errors.New("missing title or invalid year")
	ErrMovieIDRequired  = errors.New("movieID required")

	// Character domain errors
	ErrCharacterNotFound = errors.New("character not found")
)
