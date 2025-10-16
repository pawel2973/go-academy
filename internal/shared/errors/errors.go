package errors

import "errors"

var (
	ErrIDRequired      = errors.New("id required")
	ErrMovieNotFound   = errors.New("movie not found")
	ErrMovieIDRequired = errors.New("movieID required")
)
