package lasagna

// TODO: define the 'PreparationTime()' function
/* PreparationTime function return the estimate
   time for the total preparation time based on
	 the number of layers. */
func PreparationTime(layers []string, preparationTimePerLayer int) int {
	if preparationTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * preparationTimePerLayer
}

// TODO: define the 'Quantities()' function
/* Quantities function takes a slice of layers
   as parameter then determine the quantity of
	 noodles and sauce needed to make your meal. */
/* noodleQuantity variable holds the sum amount
   of noodles needed based on the number of noodles
   in the slice */
/* sauceQuantity variable holds the sum amount
   of sauce needed based on the number of sauces
   in the slice */
func Quantities(recipe []string) (int, float64) {
	var noodleQuantity int = 0
	var sauceQuantity float64 = 0
	for _, sliceElement := range recipe {
		if sliceElement == "noodles" {
			noodleQuantity += 50
		} else if sliceElement == "sauce" {
			sauceQuantity += 0.2
		}
	}
	return noodleQuantity, sauceQuantity
}

// TODO: define the 'AddSecretIngredient()' function
/* AddSecretIngredient function takes two slices and
   and writes the last element of first slice in the
	 last index of the second slice. */
/* i1 holds the last index value of first slice. */
/* i2 holds the last index value of second slice. */
func AddSecretIngredient(ingredient1 []string, ingredient2 []string) {
	i1 := len(ingredient1) - 1
	i2 := len(ingredient2) - 1
	ingredient2[i2] = ingredient1[i1]
}

// TODO: define the 'ScaleRecipe()' function
/* ScaleRecipe function takes a slice of
   quantities for 2 portions and a number
	 of portions desired return a slice of
   float64 of the amounts needed for the
	 desired number of portions. */
/* portions variable holds the float value of
   portions desired. */
/* newQuantity variable is a slice that holds
   the new quantities based on the portions. */
func ScaleRecipe(quantity []float64, numberOfPortions int) []float64 {
	portions := float64(numberOfPortions) / 2
	var newQuantity = make([]float64, len(quantity))
	for i, currentQuantity := range quantity {
		newQuantity[i] = currentQuantity * portions
	}
	return newQuantity
}
