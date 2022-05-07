package main

import "fmt"

func main() {
	input := "no"
	result := false

	switch {
	case input == "true" || input == "yes" || input == "1":
		result = true
	case input == "false" || input == "no" || input == "0":
		result = false
	default:
		result = false
	}

	fmt.Println("Input", input, "converted to boolean is", result)

}
