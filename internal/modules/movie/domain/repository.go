package domain

// MovieRepository defines the contract for movie persistence operations.
type MovieRepository interface {
	Save(Movie) Movie
	GetAll() []Movie
	Get(id string) (Movie, bool)
	Update(Movie) (Movie, error)
	Delete(id string) error
}
