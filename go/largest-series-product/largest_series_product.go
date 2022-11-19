package lsproduct

import (
	"errors"
	"sort"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	switch {
	case span < 0:
		return 0, errors.New("span must not be negative")
	case span == 0 && digits == "":
		return 1, nil
	case len(digits) < span || digits == "" && span == 1:
		return 0, errors.New("span must be smaller than string length")
	}
	var productHolder int64
	var digitCombinations []string
	var productArray []int64
	for i := range digits {
		if i+span > len(digits) {
			break
		}
		digitCombinations = append(digitCombinations, digits[i:i+span])
	}
	for _, combinationString := range digitCombinations {
		productHolder = 1
		for _, numberRune := range combinationString {
			digit, err := strconv.ParseInt(string(numberRune), 10, 64)
			if err != nil {
				return 0, errors.New("digits input must only contain digits")
			}
			productHolder *= digit
		}
		productArray = append(productArray, productHolder)
	}
	sort.Slice(productArray, func(i, j int) bool { return productArray[i] > productArray[j] })
	return productArray[0], nil

}
