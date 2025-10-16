package main

import (
	"fmt"

	"github.com/pawel2973/go-academy/configs"
	"github.com/pawel2973/go-academy/internal/app/initdata"
	repository2 "github.com/pawel2973/go-academy/internal/modules/character/repository"
	"github.com/pawel2973/go-academy/internal/modules/character/service"
	"github.com/pawel2973/go-academy/internal/modules/movie/repository"
	service2 "github.com/pawel2973/go-academy/internal/modules/movie/service"
	"github.com/pawel2973/go-academy/internal/transport/router"
	"github.com/pawel2973/go-academy/pkg/server"
)

func main() {
	cfg := configs.Load()

	// Initialize repositories and load sample data
	movieRepo := repository.NewMovieRepo()
	characterRepo := repository2.NewCharacterRepo()
	initdata.InitSample(movieRepo, characterRepo)

	// Initialize services
	movieSvc := service2.NewMovieService(movieRepo, characterRepo)
	charSvc := service.NewCharacterService(characterRepo, movieRepo)

	// Initialize server (with middleware)
	srv := server.New()

	// Initialize and register API routes
	api := router.NewAPI(movieSvc, charSvc)
	api.Register(srv)

	// Start the server on the specified address
	fmt.Printf("go-academy: env=%s addr=%s\n", cfg.Env, cfg.HTTPAddr())
	srv.Logger.Fatal(srv.Start(cfg.HTTPAddr()))
}
