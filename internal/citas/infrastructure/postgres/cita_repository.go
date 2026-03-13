package postgres

import (
	"database/sql"
	"errors"

	"aliado_ddd/internal/citas/domain"
)

type CitaRepository struct {
	db *sql.DB
}

func NewCitaRepository(db *sql.DB) *CitaRepository {
	return &CitaRepository{db: db}
}

func (r *CitaRepository) Save(cita *domain.Cita) error {
	_, err := r.db.Exec(`
		INSERT INTO citas (id, invitador_id, invitado_nombre, fecha_programada, estado)
		VALUES ($1, $2, $3, $4, $5)
	`, cita.ID(), cita.InvitadorID(), cita.InvitadoNombre(), cita.FechaProgramada(), cita.Estado())
	return err
}

func (r *CitaRepository) ByID(_ string) (*domain.Cita, error) {
	// Se deja como placeholder para mostrar contrato de repositorio en DDD.
	return nil, errors.New("ByID no implementado para demo")
}
