package main

import (
	"fmt"
)

func main() {
	message := "Hola Estación Espacial Internacional"

	for _, c := range message {
		if c != ' ' {
			newChar := c - 13

			if c >= 'a' && c <= 'z' {
				if newChar < 'a' {
					fmt.Printf("%c", newChar+26)
				} else {
					fmt.Printf("%c", newChar)
				}
			} else if c >= 'A' && c <= 'Z' {
				if newChar < 'A' {
					fmt.Printf("%c", newChar+26)
				} else {
					fmt.Printf("%c", newChar)
				}
			} else {
				fmt.Printf("%c", newChar)
			}

		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()

}
