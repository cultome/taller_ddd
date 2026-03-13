package application

import (
	"context"
	"time"

	"aliado_ddd/internal/citas/domain"
	"aliado_ddd/internal/shared/infrastructure/events"
)

// CreateCitaCommand es un DTO de aplicación.
// Se ubica fuera del dominio para no contaminar el modelo con detalles de entrada.
type CreateCitaCommand struct {
	InvitadorID     string
	InvitadoNombre  string
	FechaProgramada time.Time
}

type CreateCitaService struct {
	repo       domain.Repository
	dispatcher events.Dispatcher
}

func NewCreateCitaService(repo domain.Repository, dispatcher events.Dispatcher) *CreateCitaService {
	return &CreateCitaService{repo: repo, dispatcher: dispatcher}
}

func (s *CreateCitaService) Execute(ctx context.Context, cmd CreateCitaCommand) (string, error) {
	cita, err := domain.NewCita(cmd.InvitadorID, cmd.InvitadoNombre, cmd.FechaProgramada)
	if err != nil {
		return "", err
	}

	if err := s.repo.Save(cita); err != nil {
		return "", err
	}

	if err := s.dispatcher.Publish(ctx, cita.PullEvents()); err != nil {
		return "", err
	}

	return cita.ID(), nil
}
