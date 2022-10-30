package strand

func ToRNA(dna string) string {
	rnaMap := map[rune]rune{'A': 'U', 'G': 'C', 'C': 'G', 'T': 'A'}
	result := []rune(dna)
	for i, v := range dna {
		result[i] = rnaMap[v]
	}
	return string(result)
}
