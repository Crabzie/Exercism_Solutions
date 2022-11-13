package allergies

import (
	"strings"
)

func Allergies(allergies uint) (result []string) {
	var allergens = []string{"cats", "pollen", "chocolate", "tomatoes", "strawberries", "shellfish", "peanuts", "eggs"}
	for i, v := range []int{128, 64, 32, 16, 8, 4, 2, 1} {
		if int(allergies)&v == v {
			result = append(result, allergens[i])
		}
	}
	return
}
func AllergicTo(allergies uint, allergen string) bool {
	return strings.Contains(strings.Join(Allergies(allergies), " "), allergen)
}
