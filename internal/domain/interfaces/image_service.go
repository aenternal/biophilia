package interfaces

import "io"

type ImageService interface {
	VisualizeDistribution(aminoAcidCounts map[string]int) (io.Reader, error)
}
