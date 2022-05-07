package main

import "fmt"

func main() {
	red := -20
	convertedRed := uint8(red)

	fmt.Println("Red had a value of", red, "after converstion to uint8, it became",
		convertedRed)
}
