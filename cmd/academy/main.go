package main

import (
	"fmt"

	"github.com/pawel2973/go-academy/configs"
	"github.com/pawel2973/go-academy/internal/app/initdata"
	characterRepo "github.com/pawel2973/go-academy/internal/modules/character/repository"
	characterSvc "github.com/pawel2973/go-academy/internal/modules/character/service"
	movieRepo "github.com/pawel2973/go-academy/internal/modules/movie/repository"
	movieSvc "github.com/pawel2973/go-academy/internal/modules/movie/service"
	"github.com/pawel2973/go-academy/internal/transport/router"
	"github.com/pawel2973/go-academy/pkg/server"
)

func main() {
	cfg := configs.Load()

	// Initialize repositories and load sample data
	movieRepository := movieRepo.NewMovieRepo()
	characterRepository := characterRepo.NewCharacterRepo()
	initdata.InitSample(movieRepository, characterRepository)

	// Initialize services
	movieService := movieSvc.NewMovieService(movieRepository, characterRepository)
	characterService := characterSvc.NewCharacterService(characterRepository, movieRepository)

	// Initialize server (with middleware)
	srv := server.New()

	// Initialize and register API routes
	api := router.NewAPI(movieService, characterService)
	api.Register(srv)

	// Start the server on the specified address
	fmt.Printf("go-academy: env=%s addr=%s\n", cfg.Env, cfg.HTTPAddr())
	srv.Logger.Fatal(srv.Start(cfg.HTTPAddr()))
}
