package domain

// Character represents a character in a movie with a unique identifier, associated movie ID, and name.
type Character struct {
	ID      string `json:"id"`
	MovieID string `json:"movie_id"`
	Name    string `json:"name"`
}
