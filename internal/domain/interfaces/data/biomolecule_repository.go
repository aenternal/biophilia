package data

import (
	"biophilia/internal/domain/entities"
)

type BiomoleculeRepository interface {
	Add(biomolecule entities.AddBiomolecule) (*entities.Biomolecule, error)
	GetAll() ([]*entities.Biomolecule, error)
	GetByID(id int) (*entities.Biomolecule, error)
	Update(id int, biomolecule entities.UpdateBiomolecule) (*entities.Biomolecule, error)
	Delete(id int) error
}
