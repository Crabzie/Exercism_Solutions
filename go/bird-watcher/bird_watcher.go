package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.

/* totalBirds variable holds the sum of birds counted. */
/* bird variable holds the number of birds
depeding on the index position in the slice. */
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int = 0
	for _, bird := range birdsPerDay {
		totalBirds += bird
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.

/* weekStartIndex holds the start/begining index depending on week value. */
/* weekStartIndex holds the week end/final index depending on week value. */
/* totalBirds holds the sum of birds counted. */
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStartIndex := (week * 7) - 7
	weekEndInex := week * 7
	var totalBirds int = 0
	for i := weekStartIndex; i < weekEndInex; i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
