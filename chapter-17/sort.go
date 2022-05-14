package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	
	// even more simple sort
	sort.Strings(planets)
	fmt.Println(planets)
}