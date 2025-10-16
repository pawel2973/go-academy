package initdata

import (
	charDomain "github.com/pawel2973/go-academy/internal/modules/character/domain"
	movieDomain "github.com/pawel2973/go-academy/internal/modules/movie/domain"
)

func InitSample(movies movieDomain.MovieRepository, characters charDomain.CharacterRepository) {
	starWars := movies.Save(movieDomain.Movie{Title: "Star Wars", Year: 1977})
	lordOfTheRings := movies.Save(movieDomain.Movie{Title: "The Lord of the Rings", Year: 2001})
	gladiator := movies.Save(movieDomain.Movie{Title: "Gladiator", Year: 2000})

	characters.Save(charDomain.Character{Name: "Luke Skywalker", MovieID: starWars.ID})
	characters.Save(charDomain.Character{Name: "Princess Leia", MovieID: starWars.ID})
	characters.Save(charDomain.Character{Name: "Han Solo", MovieID: starWars.ID})

	characters.Save(charDomain.Character{Name: "Frodo Baggins", MovieID: lordOfTheRings.ID})
	characters.Save(charDomain.Character{Name: "Gandalf", MovieID: lordOfTheRings.ID})
	characters.Save(charDomain.Character{Name: "Aragorn", MovieID: lordOfTheRings.ID})

	characters.Save(charDomain.Character{Name: "Maximus", MovieID: gladiator.ID})
	characters.Save(charDomain.Character{Name: "Commodus", MovieID: gladiator.ID})
	characters.Save(charDomain.Character{Name: "Lucilla", MovieID: gladiator.ID})
}
