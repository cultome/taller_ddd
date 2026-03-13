package application

import (
	"context"

	"aliado_ddd/internal/catalogos/domain"
	"aliado_ddd/internal/shared/infrastructure/events"
)

type CreateUsuarioCommand struct {
	Nombre string
	Email  string
}

type CreateUsuarioService struct {
	repo       domain.UsuarioRepository
	dispatcher events.Dispatcher
}

func NewCreateUsuarioService(repo domain.UsuarioRepository, dispatcher events.Dispatcher) *CreateUsuarioService {
	return &CreateUsuarioService{repo: repo, dispatcher: dispatcher}
}

func (s *CreateUsuarioService) Execute(ctx context.Context, cmd CreateUsuarioCommand) (string, error) {
	usuario, err := domain.NewUsuario(cmd.Nombre, cmd.Email)
	if err != nil {
		return "", err
	}

	if err := s.repo.Save(usuario); err != nil {
		return "", err
	}

	if err := s.dispatcher.Publish(ctx, usuario.PullEvents()); err != nil {
		return "", err
	}

	return usuario.ID(), nil
}
