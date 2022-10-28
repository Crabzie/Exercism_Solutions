package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	var cleanId []int
	var sum int
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
		temp, _ := strconv.Atoi(string(char))
		cleanId = append(cleanId, temp)
	}
	for i := len(cleanId) - 2; i >= 0; i -= 2 {
		if cleanId[i] >= 5 {
			cleanId[i] = (cleanId[i] * 2) - 9
		} else {
			cleanId[i] *= 2
		}
	}
	for i := 0; i < len(cleanId); i++ {
		sum += cleanId[i]
	}
	return sum%10 == 0
}
