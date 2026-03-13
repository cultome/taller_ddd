package http

import (
	"encoding/json"
	"net/http"
	"time"

	"aliado_ddd/internal/citas/application"
)

type Handler struct {
	createCita *application.CreateCitaService
}

func NewHandler(createCita *application.CreateCitaService) *Handler {
	return &Handler{createCita: createCita}
}

type createCitaRequest struct {
	InvitadorID     string `json:"invitador_id"`
	InvitadoNombre  string `json:"invitado_nombre"`
	FechaProgramada string `json:"fecha_programada"`
}

type createCitaResponse struct {
	ID string `json:"id"`
}

func (h *Handler) CreateCita(w http.ResponseWriter, r *http.Request) {
	var req createCitaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	fecha, err := time.Parse(time.RFC3339, req.FechaProgramada)
	if err != nil {
		http.Error(w, "fecha_programada debe estar en RFC3339", http.StatusBadRequest)
		return
	}

	id, err := h.createCita.Execute(r.Context(), application.CreateCitaCommand{
		InvitadorID:     req.InvitadorID,
		InvitadoNombre:  req.InvitadoNombre,
		FechaProgramada: fecha,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(createCitaResponse{ID: id})
}
