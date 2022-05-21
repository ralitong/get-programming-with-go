package main

import "fmt"


var curiosity struct {
	lat float64
	long float64
}
func main() {
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}