package visualization

import (
	"dna-analyzer/biosynthesis"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
	"path/filepath"
)

// CountAminoAcids counts the occurrences of each amino acid in the peptide
func CountAminoAcids(peptide string) map[string]int {
	counts := make(map[string]int)
	for _, aa := range peptide {
		counts[string(aa)]++
	}
	return counts
}

// VisualizeAminoAcidDistribution creates a bar chart of amino acid frequencies and saves it to a file
func VisualizeAminoAcidDistribution(peptide string) {
	counts := CountAminoAcids(peptide)
	p := plot.New()

	p.Title.Text = "Распределение аминокислот"
	p.X.Label.Text = "Аминокислота"
	p.Y.Label.Text = "Частота"

	bars := make(plotter.Values, len(counts))
	labels := make([]string, len(counts))
	i := 0
	for aa, count := range counts {
		bars[i] = float64(count)
		labels[i] = biosynthesis.AminoAcidNames[aa]
		i++
	}

	barChart, err := plotter.NewBarChart(bars, vg.Points(20))
	if err != nil {
		log.Fatalf("Не удалось создать график: %v", err)
	}
	p.Add(barChart)
	p.NominalX(labels...)

	outputDir := "output"
	outputFile := filepath.Join(outputDir, "amino_acid_distribution.png")

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Не удалось создать директорию: %v", err)
	}

	if err := p.Save(8*vg.Inch, 4*vg.Inch, outputFile); err != nil {
		log.Fatalf("Не удалось сохранить график: %v", err)
	}

	fmt.Printf("График сохранен в файл %s\n", outputFile)
}
