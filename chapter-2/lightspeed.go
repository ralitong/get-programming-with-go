// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 5_6_000_000 // km

	fmt.Println(distance/lightSpeed, "seconds") // Prints 186 seconds
	distance = 40_1_000_000
	fmt.Println(distance/lightSpeed, "seconds") // Prints 1337 seconds

	const warpDrive = 100_800 // km/s
	const day = 24
	distance = 96_300_000 // km
	fmt.Println("Warp drive speed", distance/warpDrive/day, "days")
}

