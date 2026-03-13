package application

import (
	"context"

	"aliado_ddd/internal/misiones/domain"
	"aliado_ddd/internal/shared/infrastructure/events"
)

type CreateMisionCommand struct {
	Nombre      string
	Descripcion string
}

type CreateMisionService struct {
	repo       domain.Repository
	dispatcher events.Dispatcher
}

func NewCreateMisionService(repo domain.Repository, dispatcher events.Dispatcher) *CreateMisionService {
	return &CreateMisionService{repo: repo, dispatcher: dispatcher}
}

func (s *CreateMisionService) Execute(ctx context.Context, cmd CreateMisionCommand) (string, error) {
	mision, err := domain.NewMisionDesdeCero(cmd.Nombre, cmd.Descripcion)
	if err != nil {
		return "", err
	}

	if err := s.repo.Save(mision); err != nil {
		return "", err
	}

	if err := s.dispatcher.Publish(ctx, mision.PullEvents()); err != nil {
		return "", err
	}

	return mision.ID(), nil
}
