package main

import "fmt"

func main() {
	const distance = 236_000_000_000_000_000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const days = distance / lightSpeed / secondsPerDay
	const year = days / 365

	fmt.Println("The distance of Canis Major Dwarf is",year,"light years")
}