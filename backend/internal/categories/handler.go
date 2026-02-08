package categories

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateCategory is disabled/not part of service yet strictly,
// but we can just comment it out or leave it if we add it to service.
// For now, let's just implement List.
// The original code had CreateCategory but we are refactoring.
// If we want to keep CreateCategory we should add it to Service.
// Looking at original code: "Disabled per requirements" in main.go
// So we can probably just ignore it or remove it from handler.

func (h *Handler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.ListCategories(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}
