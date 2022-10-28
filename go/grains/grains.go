package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number >= 65 {
		return 0, errors.New("error: number out of range")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var total uint64
	for i := 0; i < 64; i++ {
		total += uint64(math.Pow(2, float64(i)))
	}
	return total
}
