package entities

type Protein struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Sequence    string   `json:"sequence"`
	Description string   `json:"description"`
	Domains     []string `json:"domains"`
	Structure   string   `json:"structure"`
}
