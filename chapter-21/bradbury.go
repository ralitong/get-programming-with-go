package main

import "fmt"


type location struct {
	lat float64
	long float64
}


func main() {
	var curiosity location
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}