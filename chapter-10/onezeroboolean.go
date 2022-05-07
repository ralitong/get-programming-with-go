package main

import (
	"fmt"
)

func main() {
	boolean := true
	converted := 0
	if boolean {
		converted = 1
	} else {
		converted = 0
	}

	fmt.Println("Boolean", boolean, "to int", converted)

}