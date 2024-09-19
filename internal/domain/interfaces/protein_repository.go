package interfaces

import "biophilia/internal/domain/entities"

type ProteinRepository interface {
	Add(protein entities.Protein) error
	GetAll() ([]entities.Protein, error)
	GetByID(id string) (*entities.Protein, error)
	Update(protein entities.Protein) error
	Delete(id string) error
}
