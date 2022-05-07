package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	aToZ := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("The string is", aToZ)
	fmt.Println("The string is", len(aToZ), "bytes")
	fmt.Println("The string has", utf8.RuneCountInString(aToZ),"runes")

	fmt.Println("The char ¿ is", len("¿"), "bytes")
}