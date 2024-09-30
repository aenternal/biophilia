package domain

import "io"

type BiomoleculeService interface {
	VisualizeAminoAcidDistribution(aminoAcidCounts map[string]int) (io.Reader, error)
	VisualizeNucleotideDistribution(nucleotideCounts map[string]int) (io.Reader, error)
}
