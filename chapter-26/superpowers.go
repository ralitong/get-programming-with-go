package main

import "fmt"

func main() {
	superpowers := &[3]string{"flight", "invisibility", "superstrength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
}