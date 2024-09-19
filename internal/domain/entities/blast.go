package entities

import "reflect"

type SearchRequest struct {
	Sequence string `json:"sequence"`
}

type BlastRequest struct {
	Program  string `json:"program"`
	Database string `json:"database"`
	Sequence string `json:"sequence"`
	SType    string `json:"stype"`
	Email    string `json:"email"`
}

func NewBlastRequest(sequence string) BlastRequest {
	return BlastRequest{
		Program:  "blastn",
		Database: "ipdmhcgen",
		Sequence: sequence,
		SType:    "dna",
		Email:    "mhc@alleles.org",
	}
}

func (request *BlastRequest) Map() map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(request)

	if val.Kind() == reflect.Struct {

		for i := 0; i < val.NumField(); i++ {
			fieldName := val.Type().Field(i).Name
			fieldValue := val.Field(i).String()
			result[fieldName] = fieldValue
		}
	}

	return result
}
