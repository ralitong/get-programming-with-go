package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

type getRowFn func(step int) (string, string)

func getCelsiusFarenheitRow(step int) (string, string) {
	celsius := celsius(step)
	return fmt.Sprintf("%.2f", celsius), fmt.Sprintf("%.2f", celsius.fahrenheit())
}

func getFarenheitToCelsiusRow(step int) (string, string) {
	farenheit := fahrenheit(step)
	return fmt.Sprintf("%.2f", farenheit), fmt.Sprintf("%.2f", farenheit.celsius())
}

func printRow(step int, getRow getRowFn) string {
	original, converted := getRow(step)
	return fmt.Sprintf("| %6s | %6s |", original, converted)
}

func printSeparator() {
	fmt.Println("===================")
}

func printHeader(original string, converted string) {
	printSeparator()
	fmt.Printf("| %2s     | %2s     |\n", original, converted)
}

func drawTable(getRow getRowFn) {
	printSeparator()
	for step := -40; step < 100; step = step + 5 {
		fmt.Println(printRow(step, getRow))  
	}
	printSeparator()
}

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func fahrenheitToCelsius(f fahrenheit) celsius {
	return celsius((f - 32) * (5.0 / 9.0))
}

func (k kelvin) celsius() celsius {
	return kelvinToCelsius(k)
}

func (k kelvin) fahrenheit() fahrenheit {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func (c celsius) kelvin() kelvin {
	return celsiusToKelvin(c)
}

func (c celsius) fahrenheit() fahrenheit {
	return celsiusToFahrenheit(c)
}

func (f fahrenheit) kelvin() kelvin {
	return celsiusToKelvin(fahrenheitToCelsius(f))
}

func (f fahrenheit) celsius() celsius {
	return fahrenheitToCelsius(f)
}

func main() {
	printHeader("ºC","ºF")
	drawTable(getCelsiusFarenheitRow)

	printHeader("ºF","ºC")
	drawTable(getFarenheitToCelsiusRow)
}
