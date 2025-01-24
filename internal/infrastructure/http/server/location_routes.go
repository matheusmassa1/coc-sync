package server

import (
	"coc-sync/internal/infrastructure/http/handler"

	"github.com/go-chi/chi/v5"
)

func (s *Server) setupLocationRoutes(r chi.Router) {
	locationHandler := handler.NewLocationHandler(s.app.Location, s.cfg)

	r.Route("/location", func(r chi.Router) {
		r.Get("/ingestAll", locationHandler.IngestAll)
	})
}
