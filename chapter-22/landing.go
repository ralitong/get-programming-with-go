package main

import "fmt"

type location struct {
	name    string
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lander string, lat, long coordinate) location {
	return location{lander, lat.decimal(), long.decimal()}
}


func main() {
	locations := []location{
		{name: "Columbia Memorial Station", lat: coordinate{14, 34, 6.2, 'S'}.decimal(), long: coordinate{175, 28, 21.5, 'E'}.decimal()},
		{name: "Challenger Memorial Station", lat: coordinate{1, 56, 46.3, 'S'}.decimal(), long: coordinate{354, 28, 24.2, 'E'}.decimal()},
		{name: "Bradbury Landing", lat: coordinate{4, 35, 22.2, 'S'}.decimal(), long: coordinate{137, 26, 30.1, 'E'}.decimal()},
		{name: "Elysium Planitia", lat: coordinate{4, 30, 0.2, 'N'}.decimal(), long: coordinate{135, 54, 0, 'E'}.decimal()},
	}

	for _, location := range locations {
		fmt.Printf("%+v\n",location)
	}


}
