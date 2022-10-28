// package raindrops gives the utility to generate strings based on numbers.
package raindrops

import (
	"strconv"
)

// Convert function takes the number of raindrops and returns a string based on them.
// result variable holds a specific string based on number of raindrops.
func Convert(number int) string {
	var result string
	switch {
	case number%3 == 0 && number%5 == 0 && number%7 == 0:
		result = "PlingPlangPlong"
	case number%3 == 0 && number%5 == 0:
		result = "PlingPlang"
	case number%3 == 0 && number%7 == 0:
		result = "PlingPlong"
	case number%5 == 0 && number%7 == 0:
		result = "PlangPlong"
	case number%3 == 0:
		result = "Pling"
	case number%5 == 0:
		result = "Plang"
	case number%7 == 0:
		result = "Plong"

	default:
		result = strconv.Itoa(number)
	}
	return result
}
