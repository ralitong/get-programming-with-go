package main

import (
	"fmt"
	"math"
)

func main() {
	v := 1000

	if v >= 0 && v <= math.MaxUint8 {
		fmt.Println("v fits in int8")
	} else {
		fmt.Println("v overflows in int8")
	}
}
