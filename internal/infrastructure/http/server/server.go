package server

import (
	"coc-sync/internal/app"
	"coc-sync/internal/config"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	app    *app.App
	cfg    *config.Config
}

func NewServer(app *app.App, cfg *config.Config) *Server {
	s := &Server{
		router: chi.NewRouter(),
		app:    app,
		cfg:    cfg,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
}

func (s *Server) setupRoutes() {
	s.router.Route("/api/v1", func(r chi.Router) {
		s.setupLocationRoutes(r)
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(addr, s)
}
