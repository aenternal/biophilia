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
	ID          int
	Type        BiomoleculeType
	Name        string
	Sequence    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AddBiomolecule struct {
	Name        string
	Type        BiomoleculeType
	Sequence    string
	Description string
}

type UpdateBiomolecule struct {
	Name        string
	Sequence    string
	Description string
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

func NucleotideNames() map[string]string {
	return map[string]string{"A": "Аденин", "G": "Гуанин", "C": "Цитозин", "U": "Урацил", "T": "Тимин"}
}
