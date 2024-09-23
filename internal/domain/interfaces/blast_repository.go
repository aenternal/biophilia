package interfaces

import (
	"biophilia/internal/domain/entities"
)

type BlastRepository interface {
	Search(sequence entities.SearchRequest) (string, error)
	GetSearchResults(jobId string)
}
