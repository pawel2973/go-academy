package domain

// Movie represents a film with a unique identifier, title, and release year.
type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
