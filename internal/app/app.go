package app

import (
	"context"
	"fmt"

	"github.com/pawel2973/go-academy/internal/app/initdata"
	"go.uber.org/fx"

	"github.com/pawel2973/go-academy/configs"
	"github.com/pawel2973/go-academy/internal/modules/character"
	"github.com/pawel2973/go-academy/internal/modules/movie"
	"github.com/pawel2973/go-academy/internal/transport/router"
	"github.com/pawel2973/go-academy/pkg/server"
)

// New builds the full FX app composed of all modules.
func New() *fx.App {
	return fx.New(
		// --- Global providers ---
		fx.Provide(configs.Load),

		// --- Domain modules ---
		movie.Module,
		character.Module,

		// --- Infrastructure modules ---
		router.Module,
		server.Module,

		// --- App lifecycle ---
		fx.Invoke(initdata.InitSample),
		fx.Invoke(registerHTTP),
	)
}

// registerHTTP handles HTTP routing and lifecycle hooks.
func registerHTTP(
	lc fx.Lifecycle,
	api *router.API,
	srv *server.Server,
	cfg configs.Config,
) {
	api.Register(srv.E)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.Start(cfg.HTTPAddr()); err != nil {
					fmt.Println("Server stopped:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Graceful shutdown...")
			return srv.Shutdown(ctx)
		},
	})
}
