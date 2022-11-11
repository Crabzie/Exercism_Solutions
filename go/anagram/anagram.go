// BenchmarkDetectAnagrams-2          20791             69410 ns/op            8288 B/op        205 allocs/op
package anagram

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Detect(subject string, candidates []string) (result []string) {
	word := SortString(strings.ToLower(subject))
	for _, v := range candidates {
		if strings.EqualFold(v, subject) {
			continue
		}
		candidateWord := SortString(strings.ToLower(v))
		ok := strings.EqualFold(candidateWord, word)
		if ok {
			result = append(result, v)
		}
	}
	return
}
