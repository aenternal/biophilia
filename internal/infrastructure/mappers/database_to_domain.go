package mappers

import (
	"biophilia/internal/data/repositories/database/models"
	"biophilia/internal/domain/entities"
)

func MapDatabaseBiomoleculeToDomain(biomolecule models.Biomolecule) *entities.Biomolecule {
	return &entities.Biomolecule{
		ID:          biomolecule.ID,
		Type:        biomolecule.Type,
		Name:        biomolecule.Name,
		Sequence:    biomolecule.Sequence,
		Description: biomolecule.Description,
		CreatedAt:   biomolecule.CreatedAt,
		UpdatedAt:   biomolecule.UpdatedAt,
	}
}

func MapDatabaseBiomoleculesToDomain(biomolecules []models.Biomolecule) []*entities.Biomolecule {
	var domainBiomolecules []*entities.Biomolecule
	for _, biomole := range biomolecules {
		domainBiomolecules = append(domainBiomolecules, MapDatabaseBiomoleculeToDomain(biomole))
	}
	return domainBiomolecules
}
