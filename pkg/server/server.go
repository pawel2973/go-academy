// Package server provides HTTP server functionality using the Echo framework.
package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server wraps an Echo instance to provide HTTP server capabilities.
type Server struct {
	E *echo.Echo
}

// New creates and configures a new Server with default middleware.
func New() *Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	return &Server{E: e}
}

// Start launches the HTTP server on the specified address.
func (s *Server) Start(addr string) error {
	fmt.Println("Server running on", addr)
	return s.E.Start(addr)
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down server...")
	return s.E.Shutdown(ctx)
}
