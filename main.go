package main

import (
	"dna-analyzer/biosynthesis"
	"dna-analyzer/blast"
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

	filename := fmt.Sprintf("%s_amino_acid_distribution.png", dna)
	visualization.VisualizeAminoAcidDistribution(filename, peptide)

	blastHits, err := blast.PerformEBIBLAST(dna)
	if err != nil {
		log.Fatalf("Ошибка выполнения BLAST: %v", err)
	}

	fmt.Println("Результаты BLAST:")
	blast.PrintEBIBlastHits(blastHits)

	filename = fmt.Sprintf("%s_blast_results.txt", dna)
	if err := blast.SaveResultsToFile(blastHits, filename); err != nil {
		fmt.Printf("ошибка сохранения результатов: %v\n", err)
		return
	}
}
