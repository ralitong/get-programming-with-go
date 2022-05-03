package main

import "fmt"

func main() {
	c := 'a'
	c = c + 3
	fmt.Printf("%c\n", c)
	
	c = 'g'
	c = c - 'a' + 'A'
	fmt.Printf("%c\n", c)

}