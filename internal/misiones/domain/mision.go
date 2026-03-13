package domain

import (
	"errors"
	"strings"
	"time"

	shared "aliado_ddd/internal/shared/domain"
)

type Mision struct {
	id            string
	nombre        string
	descripcion   string
	estado        string
	pendingEvents []shared.DomainEvent
}

const (
	EstadoMisionCreada   = "CREADA"
	EstadoMisionIniciada = "INICIADA"
)

func NewMisionDesdeCero(nombre, descripcion string) (*Mision, error) {
	if strings.TrimSpace(nombre) == "" {
		return nil, errors.New("nombre es requerido")
	}

	m := &Mision{
		id:          shared.NewID(),
		nombre:      nombre,
		descripcion: descripcion,
		estado:      EstadoMisionCreada,
	}

	m.raise(MisionCreadaDesdeCero{
		MisionID: m.id,
		Nombre:   m.nombre,
		When:     time.Now(),
	})

	return m, nil
}

func (m *Mision) ID() string          { return m.id }
func (m *Mision) Nombre() string      { return m.nombre }
func (m *Mision) Descripcion() string { return m.descripcion }
func (m *Mision) Estado() string      { return m.estado }

func (m *Mision) PullEvents() []shared.DomainEvent {
	evts := m.pendingEvents
	m.pendingEvents = nil
	return evts
}

func (m *Mision) raise(evt shared.DomainEvent) {
	m.pendingEvents = append(m.pendingEvents, evt)
}

type MisionCreadaDesdeCero struct {
	MisionID string
	Nombre   string
	When     time.Time
}

func (e MisionCreadaDesdeCero) EventName() string     { return "MisionCreadaDesdeCero" }
func (e MisionCreadaDesdeCero) OccurredAt() time.Time { return e.When }
func (e MisionCreadaDesdeCero) Payload() map[string]any {
	return map[string]any{"mision_id": e.MisionID, "nombre": e.Nombre}
}

// Eventos definidos (no exhaustivos) para mapear el lenguaje ubicuo del contexto.
type MisionCreadaDesdePlantilla struct {
	MisionID string
	When     time.Time
}
type MisionCreadaDesdeCargaMasiva struct {
	MisionID string
	When     time.Time
}
type MisionCreadaDesdeOTM struct {
	MisionID string
	When     time.Time
}
type EventoBitacoraMisionCreado struct {
	MisionID string
	Mensaje  string
	When     time.Time
}
type MisionIniciada struct {
	MisionID string
	When     time.Time
}
type DocumentoAMisionSubido struct {
	MisionID string
	Nombre   string
	When     time.Time
}
type MisionEditada struct {
	MisionID string
	When     time.Time
}

func (e MisionCreadaDesdePlantilla) EventName() string     { return "MisionCreadaDesdePlantilla" }
func (e MisionCreadaDesdePlantilla) OccurredAt() time.Time { return e.When }
func (e MisionCreadaDesdePlantilla) Payload() map[string]any {
	return map[string]any{"mision_id": e.MisionID}
}
func (e MisionCreadaDesdeCargaMasiva) EventName() string     { return "MisionCreadaDesdeCargaMasiva" }
func (e MisionCreadaDesdeCargaMasiva) OccurredAt() time.Time { return e.When }
func (e MisionCreadaDesdeCargaMasiva) Payload() map[string]any {
	return map[string]any{"mision_id": e.MisionID}
}
func (e MisionCreadaDesdeOTM) EventName() string     { return "MisionCreadaDesdeOTM" }
func (e MisionCreadaDesdeOTM) OccurredAt() time.Time { return e.When }
func (e MisionCreadaDesdeOTM) Payload() map[string]any {
	return map[string]any{"mision_id": e.MisionID}
}
func (e EventoBitacoraMisionCreado) EventName() string     { return "EventoBitacoraMisionCreado" }
func (e EventoBitacoraMisionCreado) OccurredAt() time.Time { return e.When }
func (e EventoBitacoraMisionCreado) Payload() map[string]any {
	return map[string]any{"mision_id": e.MisionID, "mensaje": e.Mensaje}
}
func (e MisionIniciada) EventName() string             { return "MisionIniciada" }
func (e MisionIniciada) OccurredAt() time.Time         { return e.When }
func (e MisionIniciada) Payload() map[string]any       { return map[string]any{"mision_id": e.MisionID} }
func (e DocumentoAMisionSubido) EventName() string     { return "DocumentoAMisionSubido" }
func (e DocumentoAMisionSubido) OccurredAt() time.Time { return e.When }
func (e DocumentoAMisionSubido) Payload() map[string]any {
	return map[string]any{"mision_id": e.MisionID, "nombre": e.Nombre}
}
func (e MisionEditada) EventName() string       { return "MisionEditada" }
func (e MisionEditada) OccurredAt() time.Time   { return e.When }
func (e MisionEditada) Payload() map[string]any { return map[string]any{"mision_id": e.MisionID} }

type Repository interface {
	Save(mision *Mision) error
	ByID(id string) (*Mision, error)
}
