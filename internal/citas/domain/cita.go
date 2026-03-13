package domain

import (
	"errors"
	"strings"
	"time"

	shared "aliado_ddd/internal/shared/domain"
)

// Cita es el Aggregate Root del contexto Citas.
// Toda modificación relevante debería pasar por este agregado
// para proteger invariantes de negocio.
type Cita struct {
	id              string
	invitadorID     string
	invitadoNombre  string
	fechaProgramada time.Time
	estado          string
	pendingEvents   []shared.DomainEvent
}

const (
	EstadoCitaCreada = "CREADA"
)

func NewCita(invitadorID, invitadoNombre string, fechaProgramada time.Time) (*Cita, error) {
	if strings.TrimSpace(invitadorID) == "" {
		return nil, errors.New("invitadorID es requerido")
	}
	if strings.TrimSpace(invitadoNombre) == "" {
		return nil, errors.New("invitadoNombre es requerido")
	}
	if fechaProgramada.IsZero() {
		return nil, errors.New("fechaProgramada es requerida")
	}

	c := &Cita{
		id:              shared.NewID(),
		invitadorID:     invitadorID,
		invitadoNombre:  invitadoNombre,
		fechaProgramada: fechaProgramada,
		estado:          EstadoCitaCreada,
	}

	c.raise(CitaCreada{
		ID:              c.id,
		InvitadorID:     c.invitadorID,
		InvitadoNombre:  c.invitadoNombre,
		FechaProgramada: c.fechaProgramada,
		When:            time.Now(),
	})

	return c, nil
}

func (c *Cita) ID() string                 { return c.id }
func (c *Cita) InvitadorID() string        { return c.invitadorID }
func (c *Cita) InvitadoNombre() string     { return c.invitadoNombre }
func (c *Cita) FechaProgramada() time.Time { return c.fechaProgramada }
func (c *Cita) Estado() string             { return c.estado }

func (c *Cita) PullEvents() []shared.DomainEvent {
	evts := c.pendingEvents
	c.pendingEvents = nil
	return evts
}

func (c *Cita) raise(evt shared.DomainEvent) {
	c.pendingEvents = append(c.pendingEvents, evt)
}

// CitaCreada mapea uno de los eventos
type CitaCreada struct {
	ID              string
	InvitadorID     string
	InvitadoNombre  string
	FechaProgramada time.Time
	When            time.Time
}

func (e CitaCreada) EventName() string     { return "CitaCreada" }
func (e CitaCreada) OccurredAt() time.Time { return e.When }
func (e CitaCreada) Payload() map[string]any {
	return map[string]any{
		"id":               e.ID,
		"invitador_id":     e.InvitadorID,
		"invitado_nombre":  e.InvitadoNombre,
		"fecha_programada": e.FechaProgramada,
	}
}

// Los siguientes eventos se dejan definidos para evolución posterior,
// aunque no se implementa todavía su flujo completo.
type DocumentoParaCitasSubido struct {
	CitaID string
	Nombre string
	When   time.Time
}

func (e DocumentoParaCitasSubido) EventName() string     { return "DocumentoParaCitasSubido" }
func (e DocumentoParaCitasSubido) OccurredAt() time.Time { return e.When }
func (e DocumentoParaCitasSubido) Payload() map[string]any {
	return map[string]any{"cita_id": e.CitaID, "nombre": e.Nombre}
}

type DocumentoParaCitasRevisado struct {
	CitaID  string
	Revisor string
	When    time.Time
}

func (e DocumentoParaCitasRevisado) EventName() string     { return "DocumentoParaCitasRevisado" }
func (e DocumentoParaCitasRevisado) OccurredAt() time.Time { return e.When }
func (e DocumentoParaCitasRevisado) Payload() map[string]any {
	return map[string]any{"cita_id": e.CitaID, "revisor": e.Revisor}
}

// Repository representa el puerto del dominio hacia persistencia.
// El dominio no conoce SQL, ORM ni detalles de base de datos.
type Repository interface {
	Save(cita *Cita) error
	ByID(id string) (*Cita, error)
}
