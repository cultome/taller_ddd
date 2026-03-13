package http

import (
	"encoding/json"
	"net/http"

	"aliado_ddd/internal/catalogos/application"
)

type Handler struct {
	createUsuario *application.CreateUsuarioService
}

func NewHandler(createUsuario *application.CreateUsuarioService) *Handler {
	return &Handler{createUsuario: createUsuario}
}

type createUsuarioRequest struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

type createUsuarioResponse struct {
	ID string `json:"id"`
}

func (h *Handler) CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var req createUsuarioRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	id, err := h.createUsuario.Execute(r.Context(), application.CreateUsuarioCommand{
		Nombre: req.Nombre,
		Email:  req.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(createUsuarioResponse{ID: id})
}
