package main

import (
	"fmt"

	"github.com/pawel2973/go-academy/configs"
	"github.com/pawel2973/go-academy/internal/initdata"
	"github.com/pawel2973/go-academy/internal/repository"
	"github.com/pawel2973/go-academy/internal/service"
)

func main() {
	cfg := configs.Load()

	// Initialize repositories and load sample data
	movieRepo := repository.NewMovieRepo()
	characterRepo := repository.NewCharacterRepo()
	initdata.InitSample(movieRepo, characterRepo)

	// Initialize services
	movieSvc := service.NewMovieService(movieRepo, characterRepo)
	charSvc := service.NewCharacterService(characterRepo, movieRepo)
	_, _ = movieSvc, charSvc

	fmt.Printf("go-academy: env=%s addr=%s\n", cfg.Env, cfg.HTTPAddr())
}
