package mappers

import (
	"biophilia/internal/data/repositories/database/models"
	"biophilia/internal/domain/entities"
)

func MapDomainBiomoleculeToDatabase(biomolecule entities.Biomolecule) models.Biomolecule {
	return models.Biomolecule{
		ID:          biomolecule.ID,
		Type:        biomolecule.Type,
		Name:        biomolecule.Name,
		Sequence:    biomolecule.Sequence,
		Description: biomolecule.Description,
		CreatedAt:   biomolecule.CreatedAt,
		UpdatedAt:   biomolecule.UpdatedAt,
	}
}
