package biosynthesis

import (
	"strings"
)

// CodonTable maps RNA codons to amino acids
var CodonTable = map[string]string{
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

var AminoAcidNames = map[string]string{
	"F": "Фенилаланин", "L": "Лейцин", "I": "Изолейцин", "M": "Метионин",
	"V": "Валин", "S": "Серин", "P": "Пролин", "T": "Треонин",
	"A": "Аланин", "Y": "Тирозин", "*": "Стоп", "H": "Гистидин",
	"Q": "Глутамин", "N": "Аспарагин", "K": "Лизин", "D": "Аспартат",
	"E": "Глутамат", "C": "Цистеин", "W": "Триптофан", "R": "Аргинин",
	"G": "Глицин",
}

// Translate converts mRNA to peptide sequence using the codon table
func Translate(mrna string) string {
	var peptide strings.Builder
	for i := 0; i < len(mrna)-2; i += 3 {
		codon := mrna[i : i+3]
		if aminoAcid, ok := CodonTable[codon]; ok {
			if aminoAcid == "*" {
				break
			}
			peptide.WriteString(aminoAcid)
		}
	}
	return peptide.String()
}
