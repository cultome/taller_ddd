package events

import (
	"context"
	"log"
	"time"

	"aliado_ddd/internal/shared/domain"
)

// Dispatcher define el contrato para publicar eventos.
// Se mantiene en infraestructura para poder reemplazarlo por Kafka/RabbitMQ
// sin tocar application/domain.
type Dispatcher interface {
	Publish(ctx context.Context, evts []domain.DomainEvent) error
}

// InMemoryDispatcher es una implementación mínima:
// registra eventos en logs para facilitar la lectura en clases.
type InMemoryDispatcher struct{}

func NewInMemoryDispatcher() *InMemoryDispatcher {
	return &InMemoryDispatcher{}
}

func (d *InMemoryDispatcher) Publish(_ context.Context, evts []domain.DomainEvent) error {
	for _, evt := range evts {
		log.Printf("domain_event name=%s at=%s payload=%v", evt.EventName(), evt.OccurredAt().Format(time.RFC3339), evt.Payload())
	}
	return nil
}
