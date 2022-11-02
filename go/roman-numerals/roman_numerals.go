package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("invalid input")
	}
	romans := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, roman := range romans {
		for input >= roman.value {
			result.WriteString(roman.digit)
			input -= roman.value
		}
	}
	return result.String(), nil
}
