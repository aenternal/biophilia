package main

import (
	"dna-analyzer/biosynthesis"
	"dna-analyzer/utils"
	"dna-analyzer/visualization"
	"fmt"
	"log"
)

func main() {
	dna, err := utils.ReadDNASequence()
	if err != nil {
		log.Fatalf("Ошибка ввода последовательности ДНК: %v", err)
	}

	mrna := biosynthesis.Transcribe(dna)
	peptide := biosynthesis.Translate(mrna)

	fmt.Printf("мРНК: %s\n", mrna)
	fmt.Printf("Пептид: %s\n", peptide)

	visualization.VisualizeAminoAcidDistribution(peptide)
}
