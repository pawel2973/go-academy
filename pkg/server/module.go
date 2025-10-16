package server

import "go.uber.org/fx"

// Module exposes Server to the FX dependency container.
var Module = fx.Options(
	fx.Provide(New),
)
