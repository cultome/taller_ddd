package domain

import (
	"errors"
	"strings"
	"time"

	shared "aliado_ddd/internal/shared/domain"
)

// Usuario es aggregate root de un subconjunto del contexto Catálogos.
type Usuario struct {
	id            string
	nombre        string
	email         string
	permisos      []string
	pendingEvents []shared.DomainEvent
}

func NewUsuario(nombre, email string) (*Usuario, error) {
	if strings.TrimSpace(nombre) == "" {
		return nil, errors.New("nombre es requerido")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email es requerido")
	}

	u := &Usuario{
		id:       shared.NewID(),
		nombre:   nombre,
		email:    email,
		permisos: []string{},
	}
	u.raise(UsuarioCreado{UsuarioID: u.id, Nombre: u.nombre, Email: u.email, When: time.Now()})

	return u, nil
}

func (u *Usuario) ID() string         { return u.id }
func (u *Usuario) Nombre() string     { return u.nombre }
func (u *Usuario) Email() string      { return u.email }
func (u *Usuario) Permisos() []string { return append([]string{}, u.permisos...) }

func (u *Usuario) PullEvents() []shared.DomainEvent {
	evts := u.pendingEvents
	u.pendingEvents = nil
	return evts
}

func (u *Usuario) raise(evt shared.DomainEvent) {
	u.pendingEvents = append(u.pendingEvents, evt)
}

type UsuarioCreado struct {
	UsuarioID string
	Nombre    string
	Email     string
	When      time.Time
}

func (e UsuarioCreado) EventName() string     { return "UsuarioCreado" }
func (e UsuarioCreado) OccurredAt() time.Time { return e.When }
func (e UsuarioCreado) Payload() map[string]any {
	return map[string]any{"usuario_id": e.UsuarioID, "nombre": e.Nombre, "email": e.Email}
}

// Eventos adicionales del contexto definidos para evolución.
type ActivoCreado struct {
	ActivoID string
	When     time.Time
}
type UbicacionCreada struct {
	UbicacionID string
	When        time.Time
}
type DocumentoSubidoAUsuario struct {
	UsuarioID string
	Nombre    string
	When      time.Time
}
type DocumentoSubidoAActivo struct {
	ActivoID string
	Nombre   string
	When     time.Time
}
type DocumentoSubidoAUbicacion struct {
	UbicacionID string
	Nombre      string
	When        time.Time
}
type PermisosUsuarioEditados struct {
	UsuarioID string
	When      time.Time
}

func (e ActivoCreado) EventName() string        { return "ActivoCreado" }
func (e ActivoCreado) OccurredAt() time.Time    { return e.When }
func (e ActivoCreado) Payload() map[string]any  { return map[string]any{"activo_id": e.ActivoID} }
func (e UbicacionCreada) EventName() string     { return "UbicacionCreada" }
func (e UbicacionCreada) OccurredAt() time.Time { return e.When }
func (e UbicacionCreada) Payload() map[string]any {
	return map[string]any{"ubicacion_id": e.UbicacionID}
}
func (e DocumentoSubidoAUsuario) EventName() string     { return "DocumentoSubidoAUsuario" }
func (e DocumentoSubidoAUsuario) OccurredAt() time.Time { return e.When }
func (e DocumentoSubidoAUsuario) Payload() map[string]any {
	return map[string]any{"usuario_id": e.UsuarioID, "nombre": e.Nombre}
}
func (e DocumentoSubidoAActivo) EventName() string     { return "DocumentoSubidoAActivo" }
func (e DocumentoSubidoAActivo) OccurredAt() time.Time { return e.When }
func (e DocumentoSubidoAActivo) Payload() map[string]any {
	return map[string]any{"activo_id": e.ActivoID, "nombre": e.Nombre}
}
func (e DocumentoSubidoAUbicacion) EventName() string     { return "DocumentoSubidoAUbicacion" }
func (e DocumentoSubidoAUbicacion) OccurredAt() time.Time { return e.When }
func (e DocumentoSubidoAUbicacion) Payload() map[string]any {
	return map[string]any{"ubicacion_id": e.UbicacionID, "nombre": e.Nombre}
}
func (e PermisosUsuarioEditados) EventName() string     { return "PermisosUsuarioEditados" }
func (e PermisosUsuarioEditados) OccurredAt() time.Time { return e.When }
func (e PermisosUsuarioEditados) Payload() map[string]any {
	return map[string]any{"usuario_id": e.UsuarioID}
}

type UsuarioRepository interface {
	Save(usuario *Usuario) error
	ByID(id string) (*Usuario, error)
}
