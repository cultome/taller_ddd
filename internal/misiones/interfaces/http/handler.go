package http

import (
	"encoding/json"
	"net/http"

	"aliado_ddd/internal/misiones/application"
)

type Handler struct {
	createMision *application.CreateMisionService
}

func NewHandler(createMision *application.CreateMisionService) *Handler {
	return &Handler{createMision: createMision}
}

type createMisionRequest struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type createMisionResponse struct {
	ID string `json:"id"`
}

func (h *Handler) CreateMision(w http.ResponseWriter, r *http.Request) {
	var req createMisionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	id, err := h.createMision.Execute(r.Context(), application.CreateMisionCommand{
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(createMisionResponse{ID: id})
}
