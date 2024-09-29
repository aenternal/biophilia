package database

import (
	"biophilia/internal/domain/entities"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SQLXBiomoleculeRepository struct {
	db *sqlx.DB
}

func NewBiomoleculeRepository(db *sqlx.DB) *SQLXBiomoleculeRepository {
	return &SQLXBiomoleculeRepository{
		db: db,
	}
}

func (r *SQLXBiomoleculeRepository) Add(biomolecule entities.AddBiomoleculeRequest) error {
	query := `
		INSERT INTO biomolecules (type, name, sequence, description)
		VALUES (:type, :name, :sequence, :description)`

	_, err := r.db.NamedExec(query, biomolecule)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении биомолекулы: %w", err)
	}

	return nil
}

func (r *SQLXBiomoleculeRepository) GetAll() ([]entities.Biomolecule, error) {
	var biomolecules []entities.Biomolecule

	query := `SELECT * FROM biomolecules`

	err := r.db.Select(&biomolecules, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении всех биомолекул: %w", err)
	}

	return biomolecules, nil
}

func (r *SQLXBiomoleculeRepository) GetByID(id int) (*entities.Biomolecule, error) {
	var biomolecule entities.Biomolecule

	query := `SELECT * FROM biomolecules WHERE id = $1`

	err := r.db.Get(&biomolecule, query, id)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении биомолекулы по ID: %w", err)
	}

	return &biomolecule, nil
}

func (r *SQLXBiomoleculeRepository) Update(id int, biomolecule entities.UpdateBiomoleculeRequest) error {
	query := `
		UPDATE biomolecules
		SET type = :type, name = :name, sequence = :sequence, description = :description, updated_at = NOW()
		WHERE id = :id`

	params := map[string]interface{}{
		"id":          id,
		"name":        biomolecule.Name,
		"description": biomolecule.Description,
	}

	_, err := r.db.NamedExec(query, params)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении биомолекулы: %w", err)
	}

	return nil
}

func (r *SQLXBiomoleculeRepository) Delete(id int) error {
	query := `DELETE FROM biomolecules WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении биомолекулы: %w", err)
	}

	return nil
}
