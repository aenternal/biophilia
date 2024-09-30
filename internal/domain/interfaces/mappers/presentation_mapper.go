package mappers

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/presentation/http/rest/entities/requests"
	"biophilia/internal/presentation/http/rest/entities/responses"
)

type PresentationMapper interface {
	PresentationBiomoleculeToDomain(pBiomolecule responses.Biomolecule) entities.Biomolecule
	PresentationCreateBiomoleculeToDomain(pBiomolecule requests.AddBiomoleculeRequest) entities.AddBiomoleculeRequest
	PresentationUpdateBiomoleculeToDomain(pBiomolecule requests.UpdateBiomoleculeRequest) entities.UpdateBiomoleculeRequest
}
