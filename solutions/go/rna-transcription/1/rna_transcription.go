package rnatranscription

import (
	"strings"
)

var dna2rna = map[byte]byte{
	'A': 'U',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

func ToRNA(dna string) string {
	sb := new(strings.Builder)
	sb.Grow(len(dna))
	for i := range dna {
		sb.WriteByte(dna2rna[dna[i]])
	}
	return sb.String()
}
