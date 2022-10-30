package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
	for index, value := range in {
		for _, stringValue := range value {
			result[strings.ToLower(stringValue)] = index
		}
	}
	return result
}
