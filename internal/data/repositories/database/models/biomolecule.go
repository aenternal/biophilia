package models

import (
	"biophilia/internal/domain/entities"
	"time"
)

type Biomolecule struct {
	ID          int                      `db:"id"`
	Type        entities.BiomoleculeType `db:"type"`
	Name        string                   `db:"name"`
	Sequence    string                   `db:"sequence"`
	Description string                   `db:"description"`
	CreatedAt   time.Time                `db:"created_at"`
	UpdatedAt   time.Time                `db:"updated_at"`
}
