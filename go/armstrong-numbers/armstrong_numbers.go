// BenchmarkIsNumber-2       377844              3139 ns/op             162 B/op         38 allocs/op
package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	numb := strconv.Itoa(n)
	var result int
	for _, v := range numb {
		temp, _ := strconv.ParseFloat(string(v), 64)
		result += int(math.Pow(temp, float64(len(numb))))
	}
	return result == n
}
