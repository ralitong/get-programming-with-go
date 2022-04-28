package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const myNumber = 34

	for {
		var computerNumber = rand.Intn(100)
		fmt.Println("The computer guesses", computerNumber, "as the number")
		if computerNumber == myNumber {
			fmt.Println("The computer guessed correctly, the number is", myNumber)
			break
		}
		fmt.Println("The computer did not guess correctly, trying again ...")
	}
}