package pangram

import "strings"

func IsPangram(input string) bool {
	var letters strings.Builder
	letters.WriteString("abcdefghijklmnopqrstuvwxyz")
	for _, v := range letters.String() {
		if strings.Contains(strings.ToLower(input), (string(v))) {
			continue
		} else {
			return false
		}
	}
	return true
}
