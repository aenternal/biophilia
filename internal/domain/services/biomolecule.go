package services

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/domain/interfaces"
	"bytes"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"image/png"
	"io"
	"log/slog"
	"strings"
)

type BiomoleculeService struct {
	log                   *slog.Logger
	biomoleculeRepository interfaces.BiomoleculeRepository
	storageRepository     interfaces.StorageRepository
}

func NewBiomoleculeService(
	log *slog.Logger,
	biomoleculeRepository interfaces.BiomoleculeRepository,
	storageRepository interfaces.StorageRepository,
) *BiomoleculeService {
	return &BiomoleculeService{log: log, biomoleculeRepository: biomoleculeRepository, storageRepository: storageRepository}
}

func (service *BiomoleculeService) AddProtein(biomolecule entities.Biomolecule) error {
	return service.biomoleculeRepository.Add(biomolecule)
}

func (service *BiomoleculeService) GetProteins() ([]entities.Biomolecule, error) {
	return service.biomoleculeRepository.GetAll()
}

func (service *BiomoleculeService) GetProteinByID(id string) (*entities.Biomolecule, error) {
	return service.biomoleculeRepository.GetByID(id)
}

func (service *BiomoleculeService) UpdateProtein(id string, biomolecule entities.Biomolecule) error {
	return service.biomoleculeRepository.Update(id, biomolecule)
}

func (service *BiomoleculeService) DeleteProtein(id string) error {
	return service.biomoleculeRepository.Delete(id)
}

func (_ *BiomoleculeService) transcribe(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func (_ *BiomoleculeService) reverseTranscribe(rna string) string {
	return strings.ReplaceAll(rna, "U", "T")
}

func (_ *BiomoleculeService) translate(mrna string) string {
	var peptide strings.Builder
	for i := 0; i < len(mrna)-2; i += 3 {
		codon := mrna[i : i+3]
		if aminoAcid, ok := entities.CodonTable()[codon]; ok {
			if aminoAcid == "*" {
				break
			}
			peptide.WriteString(aminoAcid)
		}
	}
	return peptide.String()
}

func countSequenceUnits(sequence string) map[string]int {
	counts := make(map[string]int)
	for _, aa := range sequence {
		counts[string(aa)]++
	}
	return counts
}

func (service *BiomoleculeService) visualizeSequenceDistribution(sequence string) (io.Reader, error) {
	aminoAcidCounts := countSequenceUnits(sequence)

	plotChart := plot.New()
	plotChart.Title.Text = "Распределение аминокислот"
	plotChart.X.Label.Text = "Аминокислота"
	plotChart.Y.Label.Text = "Частота"

	barValues := make(plotter.Values, len(aminoAcidCounts))
	aminoAcidLabels := make([]string, len(aminoAcidCounts))
	i := 0
	for aminoAcid, count := range aminoAcidCounts {
		barValues[i] = float64(count)
		aminoAcidLabels[i] = entities.AminoAcidNames()[aminoAcid]
		i++
	}

	barChart, err := plotter.NewBarChart(barValues, vg.Points(20))
	if err != nil {
		service.log.Error("Не удалось создать график: %v", err)
		return nil, fmt.Errorf("не удалось создать график: %v", err)
	}

	plotChart.Add(barChart)
	plotChart.NominalX(aminoAcidLabels...)

	imageCanvas := vgimg.New(8*vg.Inch, 4*vg.Inch)
	drawingContext := draw.New(imageCanvas)

	plotChart.Draw(drawingContext)

	imageBuffer := new(bytes.Buffer)

	if err := png.Encode(imageBuffer, imageCanvas.Image()); err != nil {
		service.log.Error("Не удалось кодировать график в PNG: %v", err)
		return nil, fmt.Errorf("не удалось кодировать график в PNG: %v", err)
	}

	return imageBuffer, nil
}
