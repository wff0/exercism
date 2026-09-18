package proteintranslation

import (
	"errors"
)

var ErrInvalidBase = errors.ErrUnsupported
var ErrStop = errors.New("STOP")

func FromRNA(rna string) ([]string, error) {
	res := make([]string, 0)
	for i := 0; i < len(rna); i += 3 {
		acid, err := FromCodon(rna[i:min(i+3, len(rna))])
		if err != nil {
			if errors.Is(err, ErrInvalidBase) {
				return nil, err
			}
			break
		}

		res = append(res, acid)
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "STOP", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
