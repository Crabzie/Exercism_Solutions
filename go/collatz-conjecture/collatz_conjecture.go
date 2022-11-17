package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var result int
	if n <= 0 {
		return 0, errors.New("wrong input")
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		result++
	}
	return result, nil
}
