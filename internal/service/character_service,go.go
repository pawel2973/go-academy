package service

import (
	"github.com/pawel2973/go-academy/internal/domain"
	"github.com/pawel2973/go-academy/internal/repository"
)

// CharacterService handles use cases related to characters.
type CharacterService struct {
	characters *repository.CharacterRepo
	movies     *repository.MovieRepo
}

// NewCharacterService creates a new CharacterService.
func NewCharacterService(characters *repository.CharacterRepo, movies *repository.MovieRepo) *CharacterService {
	return &CharacterService{characters: characters, movies: movies}
}

// Create creates a character for an existing movie.
func (s *CharacterService) Create(movieID, name string) (domain.Character, error) {
	if err := s.validateMovieID(movieID); err != nil {
		return domain.Character{}, err
	}
	if err := s.validateMovieExists(movieID); err != nil {
		return domain.Character{}, err
	}
	return s.characters.Save(domain.Character{MovieID: movieID, Name: name}), nil
}

// ListByMovie returns characters associated with a movie.
func (s *CharacterService) ListByMovie(movieID string) []domain.Character {
	out := []domain.Character{}
	for _, ch := range s.characters.GetAll() {
		if ch.MovieID == movieID {
			out = append(out, ch)
		}
	}
	return out
}

// Get returns a character by its ID.
func (s *CharacterService) Get(id string) (domain.Character, bool) {
	return s.characters.Get(id)
}

// Update updates a character (requires ID; if MovieID changes, the movie must exist).
func (s *CharacterService) Update(c domain.Character) (domain.Character, error) {
	if err := s.validateCharacterHasID(c); err != nil {
		return domain.Character{}, err
	}
	if c.MovieID != "" {
		if err := s.validateMovieExists(c.MovieID); err != nil {
			return domain.Character{}, err
		}
	}
	return s.characters.Update(c)
}

// Delete removes a character by its ID.
func (s *CharacterService) Delete(id string) error {
	if id == "" {
		return ErrIDRequired
	}
	return s.characters.Delete(id)
}

// validateCharacterHasID checks if the character has a non-empty ID.
func (s *CharacterService) validateCharacterHasID(c domain.Character) error {
	if c.ID == "" {
		return ErrIDRequired
	}
	return nil
}

// validateMovieID checks if the movie ID is not empty.
func (s *CharacterService) validateMovieID(movieID string) error {
	if movieID == "" {
		return ErrMovieIDRequired
	}
	return nil
}

// validateMovieExists checks if the movie exists in the repository.
func (s *CharacterService) validateMovieExists(movieID string) error {
	if _, ok := s.movies.Get(movieID); !ok {
		return ErrMovieNotFound
	}
	return nil
}
