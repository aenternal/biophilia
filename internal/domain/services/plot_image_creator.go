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
)

type PlotImageService struct {
	log *slog.Logger
}

func NewPlotImageService(log *slog.Logger) interfaces.ImageService {
	return &PlotImageService{log: log}
}

func (service *PlotImageService) VisualizeDistribution(aminoAcidCounts map[string]int) (io.Reader, error) {
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
