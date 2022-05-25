package main

import "fmt"

type report struct {
	sol int
	temperature
	location
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	report := report{
		sol: 15,
		location: location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Printf("average %v degrees C\n", report.average())
	fmt.Printf("average %v degrees C\n", report.temperature.average())
}