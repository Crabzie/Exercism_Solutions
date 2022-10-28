package interest

// InterestRate returns the interest rate for the provided balance.
/* result variable holds the float32 value of interest depending on balance */
func InterestRate(balance float64) float32 {
	var result float32
	switch {
	case balance < 0:
		result = 3.213
	case balance >= 0 && balance < 1000:
		result = 0.5
	case balance >= 1000 && balance < 5000:
		result = 1.621
	case balance >= 5000:
		result = 2.475
	default:
		result = 0
	}
	return result
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
/* yearsRequired variable holds the integer value of years required to reach targeted balance */
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var yearsRequired int
	for i := 0; balance < targetBalance; i++ {
		balance += Interest(balance)
		yearsRequired++
	}
	return yearsRequired
}
