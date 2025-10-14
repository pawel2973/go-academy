package initdata

import (
	"github.com/pawel2973/go-academy/internal/domain"
	"github.com/pawel2973/go-academy/internal/repository"
)

func InitSample(movies *repository.MovieRepo, characters *repository.CharacterRepo) {
	starWars := movies.Save(domain.Movie{Title: "Star Wars", Year: 1977})
	lordOfTheRings := movies.Save(domain.Movie{Title: "The Lord of the Rings", Year: 2001})
	gladiator := movies.Save(domain.Movie{Title: "Gladiator", Year: 2000})

	characters.Save(domain.Character{Name: "Luke Skywalker", MovieID: starWars.ID})
	characters.Save(domain.Character{Name: "Princess Leia", MovieID: starWars.ID})
	characters.Save(domain.Character{Name: "Han Solo", MovieID: starWars.ID})

	characters.Save(domain.Character{Name: "Frodo Baggins", MovieID: lordOfTheRings.ID})
	characters.Save(domain.Character{Name: "Gandalf", MovieID: lordOfTheRings.ID})
	characters.Save(domain.Character{Name: "Aragorn", MovieID: lordOfTheRings.ID})

	characters.Save(domain.Character{Name: "Maximus", MovieID: gladiator.ID})
	characters.Save(domain.Character{Name: "Commodus", MovieID: gladiator.ID})
	characters.Save(domain.Character{Name: "Lucilla", MovieID: gladiator.ID})
}
