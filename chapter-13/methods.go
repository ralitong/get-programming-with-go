package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

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

// f = c * 9.0 / 5.0 + 32
// f = c * x + 32
// f = cx + 32
// -cx = -f + 32
// cx = f - 32
// (1/x)cx = (1/x)(f - 32)
// c = f - 32
// c = (f - 32) * 5.0 / 9.0

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
	var k kelvin = 294.0
	fmt.Println(k, "º K is ", k.celsius(), "º C")
	fmt.Println(k, "º K is ", k.fahrenheit(), "º F")
	var c celsius = 20.85
	fmt.Println(c, "º C is ", c.fahrenheit(), "º F")
	fmt.Println(c, "º C is ", c.kelvin(), "º K")
	var f fahrenheit = 69.53
	fmt.Println(f, "º F is ", f.celsius(), "º C")
	fmt.Println(f, "º F is ", c.kelvin(), "º K")
}
