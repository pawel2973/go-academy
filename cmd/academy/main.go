package main

import (
	"fmt"

	"github.com/pawel2973/go-academy/configs"
)

func main() {
	cfg := configs.Load()
	fmt.Printf("go-academy: env=%s addr=%s\n", cfg.Env, cfg.HTTPAddr())
}
