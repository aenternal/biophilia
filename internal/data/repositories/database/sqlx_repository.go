package database

import (
	"biophilia/internal/data/repositories/database/models"
	"biophilia/internal/domain/entities"
	"biophilia/internal/infrastructure/mappers"
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

func (r *SQLXBiomoleculeRepository) Add(biomolecule entities.AddBiomolecule) (*entities.Biomolecule, error) {
	var createdBiomolecule models.Biomolecule
	query := `
		INSERT INTO biomolecules (type, name, sequence, description)
		VALUES (:type, :name, :sequence, :description) RETURNING *`

	err := r.db.Get(
		&createdBiomolecule,
		query,
		biomolecule.Type,
		biomolecule.Name,
		biomolecule.Sequence,
		biomolecule.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при добавлении биомолекулы: %w", err)
	}

	return mappers.MapDatabaseBiomoleculeToDomain(createdBiomolecule), nil
}

func (r *SQLXBiomoleculeRepository) GetAll() ([]*entities.Biomolecule, error) {
	var biomolecules []models.Biomolecule

	query := `SELECT * FROM biomolecules`

	err := r.db.Select(&biomolecules, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении всех биомолекул: %w", err)
	}

	return mappers.MapDatabaseBiomoleculesToDomain(biomolecules), nil
}

func (r *SQLXBiomoleculeRepository) GetByID(id int) (*entities.Biomolecule, error) {
	var biomolecule models.Biomolecule

	query := `SELECT * FROM biomolecules WHERE id = $1`

	err := r.db.Get(&biomolecule, query, id)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении биомолекулы по ID: %w", err)
	}

	return mappers.MapDatabaseBiomoleculeToDomain(biomolecule), nil
}

func (r *SQLXBiomoleculeRepository) Update(id int, biomolecule entities.UpdateBiomolecule) (*entities.Biomolecule, error) {
	var updatedBiomolecule models.Biomolecule
	query := `
		UPDATE biomolecules
		SET name = :name, sequence = :sequence, description = :description
		WHERE id = :id
		RETURNING *`

	params := map[string]interface{}{
		"id":          id,
		"name":        biomolecule.Name,
		"sequence":    biomolecule.Sequence,
		"description": biomolecule.Description,
	}

	err := r.db.Get(&updatedBiomolecule, query, params)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении биомолекулы: %w", err)
	}

	return mappers.MapDatabaseBiomoleculeToDomain(updatedBiomolecule), nil
}

func (r *SQLXBiomoleculeRepository) Delete(id int) error {
	query := `DELETE FROM biomolecules WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении биомолекулы: %w", err)
	}

	return nil
}
