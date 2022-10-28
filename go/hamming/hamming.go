// package hamming provide utility to calculate hamming distance.
package hamming

import "errors"

// Distance function returns int value of hamming distance with an error value.
// count variable holds the number of unequalities between the two strings.
func Distance(a, b string) (int, error) {
	var count int
	if len(a) != len(b) {
		return 0, errors.New("the Strings length isn't equal")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
