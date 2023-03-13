// Package weather Provide Focecast.
package weather

// CurrentCondition represent Current weather condition.
var CurrentCondition string

// CurrentLocation represent Current City.
var CurrentLocation string

// Forecast returns a string representig condtion of weather in given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
