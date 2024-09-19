package entities

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
