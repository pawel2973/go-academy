package service

import (
	repository2 "github.com/pawel2973/go-academy/internal/modules/character/repository"
	"github.com/pawel2973/go-academy/internal/modules/movie/domain"
	"github.com/pawel2973/go-academy/internal/modules/movie/repository"
	"github.com/pawel2973/go-academy/internal/shared/errors"
)

// MovieService handles use cases related to moviesRepo.
type MovieService struct {
	moviesRepo     *repository.MovieRepo
	charactersRepo *repository2.CharacterRepo
}

// NewMovieService creates a new MovieService.
func NewMovieService(movies *repository.MovieRepo, characters *repository2.CharacterRepo) *MovieService {
	return &MovieService{moviesRepo: movies, charactersRepo: characters}
}

// Create adds a new movie (ID is assigned by the repository).
func (s *MovieService) Create(title string, year int) domain.Movie {
	return s.moviesRepo.Save(domain.Movie{Title: title, Year: year})
}

// List returns all moviesRepo.
func (s *MovieService) List() []domain.Movie {
	return s.moviesRepo.GetAll()
}

// Get returns a movie by its ID.
func (s *MovieService) Get(id string) (domain.Movie, bool) {
	return s.moviesRepo.Get(id)
}

// Update modifies a movie (requires ID inside the entity).
func (s *MovieService) Update(m domain.Movie) (domain.Movie, error) {
	if err := s.validateMovieHasID(m); err != nil {
		return domain.Movie{}, err
	}
	return s.moviesRepo.Update(m)
}

// Delete removes a movie and cascades deletion to its charactersRepo.
func (s *MovieService) Delete(id string) error {
	if err := s.validateID(id); err != nil {
		return err
	}
	if err := s.moviesRepo.Delete(id); err != nil {
		return err
	}
	for _, ch := range s.charactersRepo.GetAll() {
		if ch.MovieID == id {
			_ = s.charactersRepo.Delete(ch.ID)
		}
	}
	return nil
}

// validateID checks if the ID is not empty.
func (s *MovieService) validateID(id string) error {
	if id == "" {
		return errors.ErrIDRequired
	}
	return nil
}

// validateMovieHasID checks if the movie has a non-empty ID.
func (s *MovieService) validateMovieHasID(m domain.Movie) error {
	if m.ID == "" {
		return errors.ErrIDRequired
	}
	return nil
}
