package postgres

import (
	"database/sql"
	"errors"

	"aliado_ddd/internal/misiones/domain"
)

type MisionRepository struct {
	db *sql.DB
}

func NewMisionRepository(db *sql.DB) *MisionRepository {
	return &MisionRepository{db: db}
}

func (r *MisionRepository) Save(mision *domain.Mision) error {
	_, err := r.db.Exec(`
		INSERT INTO misiones (id, nombre, descripcion, estado)
		VALUES ($1, $2, $3, $4)
	`, mision.ID(), mision.Nombre(), mision.Descripcion(), mision.Estado())
	return err
}

func (r *MisionRepository) ByID(_ string) (*domain.Mision, error) {
	return nil, errors.New("ByID no implementado para demo")
}
