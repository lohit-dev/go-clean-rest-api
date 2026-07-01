package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lohit-dev/go-clean-rest-api/internal/store"
	"github.com/lohit-dev/go-clean-rest-api/pkg/respond"
	"go.uber.org/zap"
)

type Server struct {
	Port   string
	Router *chi.Mux
	Logger *zap.Logger
	Store  *store.Store
}

func New(port string, logger *zap.Logger, store *store.Store) *Server {
	s := &Server{
		Port:   port,
		Router: chi.NewRouter(),
		Logger: logger,
		Store:  store,
	}

	s.setupMiddleware()
	s.setupRoutes(store)

	return s
}

func (s *Server) setupMiddleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) setupRoutes(_ *store.Store) {
	s.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		respond.Ok(w, http.StatusOK, map[string]string{"status": "online"})
	})

}

func (s *Server) Run() error {
	s.Logger.Info("starting server", zap.String("port", s.Port))

	httpServer := &http.Server{
		Addr:              ":" + s.Port,
		Handler:           s.Router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- httpServer.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error("server failed", zap.Error(err))
			return err
		}
	case <-ctx.Done():
		s.Logger.Info("shutdown signal received")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			s.Logger.Error("graceful shutdown failed", zap.Error(err))
			return err
		}

		if err := <-errCh; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
