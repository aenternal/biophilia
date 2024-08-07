package blast

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// PerformEBIBLAST выполняет BLAST-запрос к EBI и возвращает результаты в текстовом формате
func PerformEBIBLAST(sequence string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"program":  "blastn",
			"database": "ipdmhcgen",
			"sequence": sequence,
			"stype":    "dna",
			"email":    "mhc@alleles.org",
		}).
		Post("https://www.ebi.ac.uk/Tools/services/rest/ncbiblast/run/")

	if err != nil {
		return "", fmt.Errorf("ошибка выполнения запроса: %v", err)
	}

	jobID := strings.TrimSpace(resp.String())
	if jobID == "" {
		return "", fmt.Errorf("не удалось получить job ID")
	}

	for {
		time.Sleep(10 * time.Second)
		statusResp, err := client.R().
			Get(fmt.Sprintf("https://www.ebi.ac.uk/Tools/services/rest/ncbiblast/status/%s", jobID))

		if err != nil {
			return "", fmt.Errorf("ошибка получения статуса: %v", err)
		}

		statusResponse := strings.TrimSpace(statusResp.String())
		if statusResponse == "" {
			return "", fmt.Errorf("не удалось получить статус выполнения %s", jobID)
		}

		if statusResponse == "FINISHED" {
			break
		}
	}

	resultsResp, err := client.R().
		SetHeader("Accept", "text/plain").
		Get(fmt.Sprintf("https://www.ebi.ac.uk/Tools/services/rest/ncbiblast/result/%s/out", jobID))

	if err != nil {
		return "", fmt.Errorf("ошибка получения результатов: %v", err)
	}

	return resultsResp.String(), nil
}

func SaveResultsToFile(results, filename string) error {
	outputDir := "output"
	outputFile := filepath.Join(outputDir, filename)

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Не удалось создать директорию: %v", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.WriteString(results)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	return nil
}

func PrintEBIBlastHits(results string) {
	fmt.Println(results)
}
