package allyourbase

import (
	"errors"
	"math"
)

func Reverse(input *[]int) {
	for i, j := 0, len(*input)-1; i < j; i, j = i+1, j-1 {
		(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
	}
}

func decimalToBase(inputDecimal *int, outputBase *int) (result []int) {
	for v := *inputDecimal; v > 0; v = v / *outputBase {
		result = append(result, v%(*outputBase))
	}
	Reverse(&result)
	return
}
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	switch {
	case inputBase < 2:
		return []int{0}, errors.New("input base must be >= 2")
	case outputBase < 2:
		return []int{0}, errors.New("output base must be >= 2")
	default:
		var temp int
		for i, v := range inputDigits {
			if v < 0 || v >= inputBase {
				return nil, errors.New("all digits must satisfy 0 <= d < input base")
			}
			temp += v * int((math.Pow(float64(inputBase), float64(len(inputDigits)-i-1))))
		}
		if temp == 0 {
			return []int{0}, nil
		}
		return decimalToBase(&temp, &outputBase), nil
	}
}
