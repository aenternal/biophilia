package data

import (
	"biophilia/internal/domain/entities"
)

type BiomoleculeRepository interface {
	Add(biomolecule entities.AddBiomoleculeRequest) error
	GetAll() ([]entities.Biomolecule, error)
	GetByID(id int) (*entities.Biomolecule, error)
	Update(id int, biomolecule entities.UpdateBiomoleculeRequest) error
	Delete(id int) error
}
