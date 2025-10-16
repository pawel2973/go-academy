package domain

// CharacterRepository defines the contract for character persistence operations.
type CharacterRepository interface {
	Save(Character) Character
	GetAll() []Character
	Get(id string) (Character, bool)
	Update(Character) (Character, error)
	Delete(id string) error
	DeleteByMovie(movieID string) int
}
