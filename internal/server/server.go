package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	Port   string
	Router *chi.Mux
	Logger *zap.Logger
	DB     *gorm.DB
}

func New(port string, logger *zap.Logger) *Server {
	s := &Server{
		Port:   port,
		Router: chi.NewRouter(),
		Logger: logger,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) setupRoutes() {
	s.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "online"})
	})

}

func (s *Server) Run() error {
	defer s.Logger.Sync()
	s.Logger.Info("starting server", zap.String("port", s.Port))

	if err := http.ListenAndServe(":"+s.Port, s.Router); err != nil {
		s.Logger.Error("Server Failed", zap.Error(err))
		return err
	}

	return nil
}
