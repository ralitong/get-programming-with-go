package main

import "fmt"

func main() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5

	fmt.Printf("%x %x %x\n", red, green, blue)
	
	// Hexadecimal with padding
	fmt.Printf("#%02x%02x%02x\n", red, green, blue)

}
