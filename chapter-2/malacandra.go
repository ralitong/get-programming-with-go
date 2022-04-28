package main

import (
	"fmt"
)

func main() {
	const duration = 28 // days
	const day = 24
	const distance = 56_000_000 // km

	fmt.Println("To reach Mars in",duration,"days, you need to travel",
	distance/duration*day,"km/h")



}