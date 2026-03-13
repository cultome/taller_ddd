package postgres

import (
	"database/sql"
	"errors"

	"aliado_ddd/internal/catalogos/domain"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) Save(usuario *domain.Usuario) error {
	_, err := r.db.Exec(`
		INSERT INTO catalogos_usuarios (id, nombre, email)
		VALUES ($1, $2, $3)
	`, usuario.ID(), usuario.Nombre(), usuario.Email())
	return err
}

func (r *UsuarioRepository) ByID(_ string) (*domain.Usuario, error) {
	return nil, errors.New("ByID no implementado para demo")
}
