package clients

import (
	"biophilia/internal/domain/entities"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

type EBIBlastClient struct {
}

func NewEBIBlastClient() *EBIBlastClient {
	return &EBIBlastClient{}
}

func (blast *EBIBlastClient) Search(sequence, database string, searchType entities.BiomoleculeType) (string, error) {
	client := resty.New()
	blastRequest := entities.NewBlastRequest(sequence, database, searchType)

	resp, err := client.R().
		SetFormData(blastRequest.Map()).
		Post("https://www.ebi.ac.uk/Tools/services/rest/ncbiblast/run/")

	if err != nil {
		return "", fmt.Errorf("ошибка выполнения запроса: %v", err)
	}

	jobID := strings.TrimSpace(resp.String())
	if jobID == "" {
		return "", fmt.Errorf("не удалось получить job ID")
	}
	return jobID, nil
}

func (blast *EBIBlastClient) GetSearchResults(jobID string) (string, error) {
	client := resty.New()
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
