package entities

import "time"

type BiomoleculeType string

const (
	BiomoleculeTypeDNA     = "dna"
	BiomoleculeTypeRNA     = "rna"
	BiomoleculeTypeProtein = "protein"
)

func (biomoleculeType BiomoleculeType) IsValid() bool {
	switch biomoleculeType {
	case BiomoleculeTypeDNA, BiomoleculeTypeRNA, BiomoleculeTypeProtein:
		return true
	}
	return false
}

type Biomolecule struct {
	ID          int       `db:"id" json:"id"`
	Type        string    `db:"type" json:"type"`
	Name        string    `db:"name" json:"name"`
	Sequence    string    `db:"sequence" json:"sequence"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

type AddBiomoleculeRequest struct {
	Type        BiomoleculeType `json:"type"`
	Name        string          `json:"name"`
	Sequence    string          `json:"sequence"`
	Description string          `json:"description"`
}

type UpdateBiomoleculeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CodonTable() map[string]string {
	return map[string]string{
		"UUU": "F", "UUC": "F", "UUA": "L", "UUG": "L",
		"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
		"AUU": "I", "AUC": "I", "AUA": "I", "AUG": "M",
		"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
		"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
		"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
		"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
		"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
		"UAU": "Y", "UAC": "Y", "UAA": "*", "UAG": "*",
		"CAU": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
		"AAU": "N", "AAC": "N", "AAA": "K", "AAG": "K",
		"GAU": "D", "GAC": "D", "GAA": "E", "GAG": "E",
		"UGU": "C", "UGC": "C", "UGA": "*", "UGG": "W",
		"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",
		"AGU": "S", "AGC": "S", "AGA": "R", "AGG": "R",
		"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
	}
}

func AminoAcidNames() map[string]string {
	return map[string]string{
		"F": "Фенилаланин", "L": "Лейцин", "I": "Изолейцин", "M": "Метионин",
		"V": "Валин", "S": "Серин", "P": "Пролин", "T": "Треонин",
		"A": "Аланин", "Y": "Тирозин", "*": "Стоп", "H": "Гистидин",
		"Q": "Глутамин", "N": "Аспарагин", "K": "Лизин", "D": "Аспартат",
		"E": "Глутамат", "C": "Цистеин", "W": "Триптофан", "R": "Аргинин",
		"G": "Глицин",
	}
}
