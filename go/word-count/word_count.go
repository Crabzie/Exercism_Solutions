// BenchmarkWordCount-2       69493             17031 ns/op            4616 B/op         48 allocs/op
package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(strings.ReplaceAll(phrase, ",", " "))
	result := make(Frequency)
	for _, word := range strings.Fields(phrase) {
		word = strings.Trim(word, "'.:!&@$%^")
		result[word]++
	}
	return result
}
