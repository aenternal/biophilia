package responses

import (
	"biophilia/internal/presentation/http/rest/entities"
	"time"
)

type Biomolecule struct {
	ID          int                      `json:"id"`
	Type        entities.BiomoleculeType `json:"type"`
	Name        string                   `json:"name"`
	Sequence    string                   `json:"sequence"`
	Description string                   `json:"description"`
	CreatedAt   time.Time                `json:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt"`
}
