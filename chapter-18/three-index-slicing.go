package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v, %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)
	
	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	dump("terrestrial", terrestrial)
	dump("planets", planets)
}