package router

import "go.uber.org/fx"

// Module registers HTTP API router and routes.
var Module = fx.Options(
	fx.Provide(NewAPI),
)
