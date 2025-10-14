package service

import (
	"github.com/pawel2973/go-academy/internal/domain"
	"github.com/pawel2973/go-academy/internal/repository"
)

// MovieService handles use cases related to movies.
type MovieService struct {
	movies     *repository.MovieRepo
	characters *repository.CharacterRepo
}

// NewMovieService creates a new MovieService.
func NewMovieService(movies *repository.MovieRepo, characters *repository.CharacterRepo) *MovieService {
	return &MovieService{movies: movies, characters: characters}
}

// Create adds a new movie (ID is assigned by the repository).
func (s *MovieService) Create(title string, year int) domain.Movie {
	return s.movies.Save(domain.Movie{Title: title, Year: year})
}

// List returns all movies.
func (s *MovieService) List() []domain.Movie {
	return s.movies.GetAll()
}

// Get returns a movie by its ID.
func (s *MovieService) Get(id string) (domain.Movie, bool) {
	return s.movies.Get(id)
}

// Update modifies a movie (requires ID inside the entity).
func (s *MovieService) Update(m domain.Movie) (domain.Movie, error) {
	if err := s.validateMovieHasID(m); err != nil {
		return domain.Movie{}, err
	}
	return s.movies.Update(m)
}

// Delete removes a movie and cascades deletion to its characters.
func (s *MovieService) Delete(id string) error {
	if err := s.validateID(id); err != nil {
		return err
	}
	if err := s.movies.Delete(id); err != nil {
		return err
	}
	for _, ch := range s.characters.GetAll() {
		if ch.MovieID == id {
			_ = s.characters.Delete(ch.ID)
		}
	}
	return nil
}

// validateID checks if the ID is not empty.
func (s *MovieService) validateID(id string) error {
	if id == "" {
		return ErrIDRequired
	}
	return nil
}

// validateMovieHasID checks if the movie has a non-empty ID.
func (s *MovieService) validateMovieHasID(m domain.Movie) error {
	if m.ID == "" {
		return ErrIDRequired
	}
	return nil
}
