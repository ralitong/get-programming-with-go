package main

import "fmt"

func main() {
	caesarString := "L fdph, L vdz, L frqtxhuhg."

	for _, c := range caesarString {
		if c >= 'a' && c <= 'z' {
			decrypted := c - 3
			if decrypted < 'a' {
				fmt.Printf("%c", decrypted+26)
			} else {
				fmt.Printf("%c", decrypted)
			}
		} else if c >= 'A' && c <= 'Z' {
			decrypted := c - 3
			if decrypted < 'A' {
				fmt.Printf("%c", decrypted+26)
			} else {
				fmt.Printf("%c", decrypted)
			}
		} else {
			fmt.Printf("%c", c)
		}

	}
	fmt.Println()

}
