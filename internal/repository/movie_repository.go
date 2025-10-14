package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pawel2973/go-academy/internal/domain"
)

// movieMap stores Movie objects by their ID.
type movieMap = map[string]domain.Movie

// MovieRepo manages operations on Movie entities.
type MovieRepo struct {
	data movieMap
}

// NewMovieRepo creates a new MovieRepo instance.
func NewMovieRepo() *MovieRepo {
	return &MovieRepo{data: make(movieMap)}
}

// Save adds a new Movie to the repository.
func (r *MovieRepo) Save(m domain.Movie) domain.Movie {
	m.ID = uuid.New().String()
	r.data[m.ID] = m
	return m
}

// GetAll returns all Movies from the repository.
func (r *MovieRepo) GetAll() []domain.Movie {
	list := make([]domain.Movie, 0, len(r.data))
	for _, movie := range r.data {
		list = append(list, movie)
	}
	return list
}

// Get retrieves a Movie by its ID.
func (r *MovieRepo) Get(id string) (domain.Movie, bool) {
	movie, ok := r.data[id]
	return movie, ok
}

// Update modifies an existing Movie.
func (r *MovieRepo) Update(m domain.Movie) (domain.Movie, error) {
	if _, ok := r.data[m.ID]; !ok {
		return domain.Movie{}, errors.New("movie not found")
	}
	r.data[m.ID] = m
	return m, nil
}

// Delete removes a Movie by its ID.
func (r *MovieRepo) Delete(id string) error {
	if _, ok := r.data[id]; !ok {
		return errors.New("movie not found")
	}
	delete(r.data, id)
	return nil
}
