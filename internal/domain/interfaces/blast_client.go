package interfaces

import "biophilia/internal/domain/entities"

type BlastClient interface {
	Search(sequence, database string, searchType entities.BiomoleculeType) (string, error)
	GetSearchResults(jobId string) (string, error)
}
