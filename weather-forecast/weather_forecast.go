// Package weather provides ability to print the city and weather condition.
package weather

// CurrentCondition is the name of the weather condition.
var CurrentCondition string
// CurrentLocation is the name of the city.
var CurrentLocation string

//Forecast takes in the city and condition and returns the forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
