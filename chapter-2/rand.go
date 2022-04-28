package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	var randomDistance = rand.Intn(345_000_001) + 56_000_000
	fmt.Println("Random distance is", randomDistance)

}
