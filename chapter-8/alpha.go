package main

import "fmt"

func main() {
	const lightSpeed = 299792
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is",days,"days of travel at light speed.")

	// var andromedaDistance uint64 = 24e18 -- will cause an overflow
	// fmt.Println("The distance of Andromeda Galaxy", andromedaDistance)

	
}