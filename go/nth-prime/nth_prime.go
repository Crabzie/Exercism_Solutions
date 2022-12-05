package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func checkIfPrime(n int) bool {
	for j := 2; j < n/2; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
func Nth(n int) (int, error) {
	switch n {
	case 0:
		return 0, errors.New("there is no zeroth prime")
	case 1:
		return 2, nil
	case 2:
		return 3, nil
	default:
		var result int
		for result = 4; n > 1; result++ {
			ok := checkIfPrime(result)
			if ok {
				n--
			}
		}
		return result - 1, nil
	}
}
