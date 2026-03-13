package domain

import "time"

// DomainEvent representa un hecho relevante de negocio ya ocurrido.
// Es la pieza central para enseñar event-driven dentro de DDD.
type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
	Payload() map[string]any
}
