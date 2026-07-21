// Package server provides the HTTP server lifecycle management for the ZoneBridge platform.
package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/philaturo/zonebridge-platform/internal/platform/config"
	"github.com/philaturo/zonebridge-platform/internal/platform/errors"
)

// Server wraps the HTTP server and provides controlled lifecycle management.
type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
	cfg        *config.Config
}

// New creates a new Server instance with the configured router and middleware.
func New(cfg *config.Config, logger *slog.Logger, router *chi.Mux) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort),
			Handler:      router,
			ReadTimeout:  cfg.ServerReadTimeout,
			WriteTimeout: cfg.ServerWriteTimeout,
			IdleTimeout:  60 * time.Second, // Prevents slowloris attacks
		},
		logger: logger,
		cfg:    cfg,
	}
}

// Start begins listening for HTTP requests. It is passive and blocks until 
// the server is shut down or encounters a fatal error. Orchestration belongs 
// in the application package.
func (s *Server) Start() error {
	s.logger.Info("starting http server", "addr", s.httpServer.Addr)
	
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%w: %v", errors.ErrServerStartup, err)
	}
	
	return nil
}

// Stop gracefully shuts down the HTTP server, respecting the configured shutdown timeout.
func (s *Server) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, s.cfg.ServerShutdownTimeout)
	defer cancel()

	s.logger.Info("initiating graceful shutdown")
	
	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("%w: %v", errors.ErrGracefulShutdown, err)
	}
	
	s.logger.Info("http server stopped gracefully")
	return nil
}