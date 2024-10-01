package requests

import (
	"biophilia/internal/domain/entities"
)

type AddBiomoleculeRequest struct {
	Name        string                   `json:"name" example:"Hemoglobin"`
	Type        entities.BiomoleculeType `json:"type" example:"protein"`
	Sequence    string                   `json:"sequence" example:"MVHLTPEEKSA"`
	Description string                   `json:"description" example:"Essential for oxygen presentation"`
}

type UpdateBiomoleculeRequest struct {
	Name        string `json:"name" example:"Hemoglobin"`
	Description string `json:"description" example:"Essential for oxygen presentation"`
}
