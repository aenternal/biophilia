package models

import "time"

type BiomoleculeFile struct {
	id            int       `db:"id"`
	biomoleculeId int       `db:"biomolecule_id"`
	fileName      string    `db:"file_name"`
	fileType      string    `db:"file_type"`
	uploadedAt    time.Time `db:"uploaded_at"`
}
