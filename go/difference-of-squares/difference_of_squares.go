package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		count += i
	}
	count = int(math.Pow(float64(count), float64(2)))
	return count
}

func SumOfSquares(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		count += int(math.Pow(float64(i), float64(2)))
	}
	return count
}

func Difference(n int) int {
	suOs := SumOfSquares(n)
	sqOs := SquareOfSum(n)
	result := sqOs - suOs
	return result
}
