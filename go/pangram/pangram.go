package pangram

import "strings"

func IsPangram(input string) bool {
	var letters strings.Builder
	letters.WriteString("abcdefghijklmnopqrstuvwxyz")
	for _, v := range letters.String() {
		if strings.ContainsRune(strings.ToLower(input), v) {
			continue
		} else {
			return false
		}
	}
	return true
}
