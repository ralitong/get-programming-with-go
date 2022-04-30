package main

import (
	"math/rand"
	"fmt"

) 






func main() {
	
	spaceLines := [3]string{"Virgin Galactic", "SpaceX", "Space Adventures"}
	trips := [2]string{"One-way", "Round-trip"}

	fmt.Printf("%-15v %4v %4v   %2v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")
	
	for i:= 10; i > 0; i-- {
		randomSpaceLine := spaceLines[rand.Intn(3)]
		randomDays := rand.Intn(30) + 20
		randomTrip := trips[rand.Intn(2)]
		randomPrice := rand.Intn(37) + 63
		fmt.Printf("%-16v %2v %-10v    $ %2v\n", randomSpaceLine, randomDays, randomTrip, randomPrice)
	}
}