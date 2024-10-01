package mappers

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/presentation/http/rest/entities/requests"
)

func MapCreateBiomoleculeRequestToDomain(biomolecule requests.AddBiomoleculeRequest) entities.AddBiomolecule {
	return entities.AddBiomolecule{
		Type:        biomolecule.Type,
		Name:        biomolecule.Name,
		Sequence:    biomolecule.Sequence,
		Description: biomolecule.Description,
	}
}

func MapUpdateBiomoleculeRequestToDomain(biomolecule requests.UpdateBiomoleculeRequest) entities.UpdateBiomolecule {
	return entities.UpdateBiomolecule{
		Name:        biomolecule.Name,
		Description: biomolecule.Description,
	}
}
