package utils

import (
	"fmt"
)

// ReadDNASequence reads a DNA sequence from standard input
func ReadDNASequence() (string, error) {
	var dna string
	fmt.Println("Введите последовательность ДНК:")
	_, err := fmt.Scanln(&dna)
	if err != nil {
		return "", err
	}
	return dna, nil
}
