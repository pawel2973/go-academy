package main

import (
	"fmt"

	"github.com/pawel2973/go-academy/configs"
	"github.com/pawel2973/go-academy/internal/initdata"
	"github.com/pawel2973/go-academy/internal/repository"
)

func main() {
	cfg := configs.Load()
	movieRepo := repository.NewMovieRepo()
	characterRepo := repository.NewCharacterRepo()
	initdata.InitSample(movieRepo, characterRepo)

	fmt.Printf("go-academy: env=%s addr=%s\n", cfg.Env, cfg.HTTPAddr())
}
