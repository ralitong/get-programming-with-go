package main

import "fmt"

func main() {
	age := 41
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays

	fmt.Println("I am", marsAge, "years old on Mars.")
	fmt.Println("Truncated earth days:", int(earthDays))

	// fmt.Println("age > marsAge", age > marsAge) will not work since they are different
	// types
}
