// BenchmarkCodon-2         7079599               144.3 ns/op             0 B/op          0 allocs/op
// BenchmarkProtein-2        570000              1953 ns/op             384 B/op         11 allocs/op
package protein

import (
	"errors"
)

var ErrInvalidBase, ErrStop = errors.New("error invalid base"), errors.New("error stop")

func FromRNA(rna string) (result []string, err error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	rnaString := []rune(rna)
	for i := 0; i < len(rna); i = i + 3 {
		proteinResult, err := FromCodon(string(rnaString[i : i+3]))
		if err != nil {
			return result, nil
		}
		result = append(result, proteinResult)
	}
	return
}

func FromCodon(codon string) (string, error) {
	switch v := codon; {
	case v == "AUG":
		return "Methionine", nil
	case v == "UUU" || v == "UUC":
		return "Phenylalanine", nil
	case v == "UUA" || v == "UUG":
		return "Leucine", nil
	case v == "UCU" || v == "UCC" || v == "UCA" || v == "UCG":
		return "Serine", nil
	case v == "UAU" || v == "UAC":
		return "Tyrosine", nil
	case v == "UGU" || v == "UGC":
		return "Cysteine", nil
	case v == "UGG":
		return "Tryptophan", nil
	case v == "UAA" || v == "UAG" || v == "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
