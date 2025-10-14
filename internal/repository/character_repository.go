package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pawel2973/go-academy/internal/domain"
)

// characterMap stores Character objects by their ID.
type characterMap = map[string]domain.Character

// CharacterRepo manages operations on Character entities.
type CharacterRepo struct {
	data characterMap
}

// NewCharacterRepo creates a new CharacterRepo instance.
func NewCharacterRepo() *CharacterRepo {
	return &CharacterRepo{data: make(characterMap)}
}

// Save adds a new Character to the repository.
func (r *CharacterRepo) Save(c domain.Character) domain.Character {
	c.ID = uuid.New().String()
	r.data[c.ID] = c
	return c
}

// GetAll returns all Characters from the repository.
func (r *CharacterRepo) GetAll() []domain.Character {
	list := make([]domain.Character, 0, len(r.data))
	for _, character := range r.data {
		list = append(list, character)
	}
	return list
}

// Get retrieves a Character by its ID.
func (r *CharacterRepo) Get(id string) (domain.Character, bool) {
	character, ok := r.data[id]
	return character, ok
}

// Update modifies an existing Character.
func (r *CharacterRepo) Update(c domain.Character) (domain.Character, error) {
	if _, ok := r.data[c.ID]; !ok {
		return domain.Character{}, errors.New("character not found")
	}
	r.data[c.ID] = c
	return c, nil
}

// Delete removes a Character by its ID.
func (r *CharacterRepo) Delete(id string) error {
	if _, ok := r.data[id]; !ok {
		return errors.New("character not found")
	}
	delete(r.data, id)
	return nil
}

// DeleteByMovie removes all Characters associated with a specific Movie ID.
func (r *CharacterRepo) DeleteByMovie(movieID string) int {
	count := 0
	for id, ch := range r.data {
		if ch.MovieID == movieID {
			delete(r.data, id)
			count++
		}
	}
	return count
}
