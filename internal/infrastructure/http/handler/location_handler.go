package handler

import (
	"coc-sync/internal/config"
	"coc-sync/internal/domain/location"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationHandler struct {
	BaseHandler
	service location.IService
}

func NewLocationHandler(service location.IService, cfg *config.Config) *LocationHandler {
	return &LocationHandler{
		BaseHandler: BaseHandler{cfg: cfg},
		service:     service,
	}
}

func (h *LocationHandler) IngestAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), h.cfg.HTTPTimeout)
	defer cancel()

	locations, err := h.service.GetLocations(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	response := map[string]interface{}{
		"message": "success",
		"count":   len(locations),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
