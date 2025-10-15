package main

import (
	"fmt"

	"github.com/pawel2973/go-academy/configs"
	"github.com/pawel2973/go-academy/internal/initdata"
	"github.com/pawel2973/go-academy/internal/repository"
	"github.com/pawel2973/go-academy/internal/service/character"
	"github.com/pawel2973/go-academy/internal/service/movie"
	"github.com/pawel2973/go-academy/internal/transport/api/router"
	"github.com/pawel2973/go-academy/pkg/server"
)

func main() {
	cfg := configs.Load()

	// Initialize repositories and load sample data
	movieRepo := repository.NewMovieRepo()
	characterRepo := repository.NewCharacterRepo()
	initdata.InitSample(movieRepo, characterRepo)

	// Initialize services
	movieSvc := movie.NewMovieService(movieRepo, characterRepo)
	charSvc := character.NewCharacterService(characterRepo, movieRepo)

	// Initialize server (with middleware)
	srv := server.New()

	// Initialize and register API routes
	api := router.NewAPI(movieSvc, charSvc)
	api.Register(srv)

	// Start the server on the specified address
	fmt.Printf("go-academy: env=%s addr=%s\n", cfg.Env, cfg.HTTPAddr())
	srv.Logger.Fatal(srv.Start(cfg.HTTPAddr()))
}
