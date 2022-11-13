package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var result strings.Builder
	for i, v := range s {
		if i == 0 {
			result.WriteRune(v)
			continue
		}
		if (s[i-1] == ' ' || s[i-1] == '_' || s[i-1] == '-') && unicode.IsLetter(v) {
			result.WriteString(strings.ToUpper(string(v)))
		}
	}
	return result.String()
}
