package entities

type BiomoleculeType string

const (
	BiomoleculeTypeDNA     = "dna"
	BiomoleculeTypeRNA     = "rna"
	BiomoleculeTypeProtein = "protein"
)

func (biomoleculeType BiomoleculeType) IsValid() bool {
	switch biomoleculeType {
	case BiomoleculeTypeDNA, BiomoleculeTypeRNA, BiomoleculeTypeProtein:
		return true
	}
	return false
}
