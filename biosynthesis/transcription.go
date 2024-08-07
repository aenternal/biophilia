package biosynthesis

import "strings"

// Transcribe converts DNA to RNA by replacing 'T' with 'U'
func Transcribe(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}
