package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New creates and configures an Echo instance with global middleware.
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover()) // Recover from panics and return HTTP 500 instead of crashing the server.
	e.Use(middleware.Logger())  // Log each HTTP request.

	return e
}
