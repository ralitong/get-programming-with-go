package main

import "fmt"

// location with a latitude, longitude in decimal degress.
type location struct {
	lat, long coordinate
}

// String formats a location with latitude, longitude
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct { 
	d, m, s float64
	h rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%v\" %c", c.d, c.m, c.s, c.h)
}

func main() {
	elysium := location{
		coordinate{4,30,0.0,'N'},
		coordinate{135,54,0.0,'E'},
	}

	fmt.Printf("Elysium Planitia is at %v\n", elysium)
}
