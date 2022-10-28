// Package weather provide tools to return city conditon based on location.
package weather

// CurrentCondition variable is used to hold the current condition of the city.
var CurrentCondition string
// CurrentLocation variable is used to hold the targeted location.
var CurrentLocation string
// Forecast function is used to return a string holding the location and its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
