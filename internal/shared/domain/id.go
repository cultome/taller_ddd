package domain

import "github.com/google/uuid"

// NewID genera IDs únicos para entidades/agregados.
// En un sistema real puede reemplazarse por ULID, KSUID o IDs de DB,
// pero para propósitos didácticos UUID es claro y suficiente.
func NewID() string {
	return uuid.NewString()
}
