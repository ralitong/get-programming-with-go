package main

import "fmt"

// kelvinToCelsius converts K to C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// celsiusToFahrenheit converts C to F
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0 ) + 32.0
	
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println(kelvin, "º K is", celsius, "º C")
	fmt.Println(celsius, "º C is", fahrenheit, "º F")
}