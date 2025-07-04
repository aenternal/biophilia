package services

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/domain/interfaces/domain"
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
)

type PlotImageService struct {
	log *slog.Logger
}

func NewPlotImageService(log *slog.Logger) domain.ImageService {
	return &PlotImageService{log: log}
}

func (service *PlotImageService) VisualizeAminoAcidDistribution(aminoAcidCounts map[string]int) (io.Reader, error) {
	barValues, aminoAcidLabels := service.preparePlotData(aminoAcidCounts, entities.AminoAcidNames())

	plotChart, err := service.createPlot(barValues, aminoAcidLabels, "Распределение аминокислот")
	if err != nil {
		return nil, service.logAndReturnError("Не удалось создать график", err)
	}

	imageBuffer, err := service.renderPlotToImage(plotChart)
	if err != nil {
		return nil, service.logAndReturnError("Не удалось отрисовать график", err)
	}

	return imageBuffer, nil
}

func (service *PlotImageService) VisualizeNucleotideDistribution(nucleotideCounts map[string]int) (io.Reader, error) {
	barValues, nucleotideLabels := service.preparePlotData(nucleotideCounts, entities.NucleotideNames())

	plotChart, err := service.createPlot(barValues, nucleotideLabels, "Распределение нуклеотидов")
	if err != nil {
		return nil, service.logAndReturnError("Не удалось создать график", err)
	}

	imageBuffer, err := service.renderPlotToImage(plotChart)
	if err != nil {
		return nil, service.logAndReturnError("Не удалось отрисовать график", err)
	}

	return imageBuffer, nil
}

func (service *PlotImageService) preparePlotData(counts map[string]int, names map[string]string) (plotter.Values, []string) {
	barValues := make(plotter.Values, len(counts))
	labels := make([]string, len(counts))

	i := 0
	for key, count := range counts {
		barValues[i] = float64(count)
		labels[i] = names[key]
		i++
	}

	return barValues, labels
}

func (service *PlotImageService) createPlot(barValues plotter.Values, labels []string, title string) (*plot.Plot, error) {
	plotChart := plot.New()
	plotChart.Title.Text = title
	plotChart.X.Label.Text = "Название"
	plotChart.Y.Label.Text = "Частота"

	barChart, err := plotter.NewBarChart(barValues, vg.Points(20))
	if err != nil {
		return nil, err
	}

	plotChart.Add(barChart)
	plotChart.NominalX(labels...)

	return plotChart, nil
}

func (service *PlotImageService) renderPlotToImage(plotChart *plot.Plot) (*bytes.Buffer, error) {
	imageCanvas := vgimg.New(8*vg.Inch, 4*vg.Inch)
	drawingContext := draw.New(imageCanvas)

	plotChart.Draw(drawingContext)

	imageBuffer := new(bytes.Buffer)
	if err := png.Encode(imageBuffer, imageCanvas.Image()); err != nil {
		return nil, err
	}

	return imageBuffer, nil
}

func (service *PlotImageService) logAndReturnError(message string, err error) error {
	service.log.Error(fmt.Sprintf("%s: %v", message, err))
	return fmt.Errorf("%s: %v", message, err)
}
