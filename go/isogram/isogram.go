package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	s := strings.ToLower(word)
	for _, char := range s {
		if !unicode.IsLetter(char) {
			continue
		}
		if charCount := strings.Count(s, string(char)); charCount > 1 {
			return false
		}
	}
	return true
}
