package mappers

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/presentation/http/rest/entities/responses"
)

func MapDomainBiomoleculeToResponse(biomolecule entities.Biomolecule) responses.Biomolecule {
	return responses.Biomolecule{
		ID:          biomolecule.ID,
		Type:        biomolecule.Type,
		Name:        biomolecule.Name,
		Sequence:    biomolecule.Sequence,
		Description: biomolecule.Description,
		CreatedAt:   biomolecule.CreatedAt,
		UpdatedAt:   biomolecule.UpdatedAt,
	}
}

func MapDomainBiomoleculesToResponse(biomolecules []entities.Biomolecule) []responses.Biomolecule {
	var responseBiomolecules []responses.Biomolecule
	for _, biomole := range biomolecules {
		responseBiomolecules = append(responseBiomolecules, MapDomainBiomoleculeToResponse(biomole))
	}
	return responseBiomolecules
}
