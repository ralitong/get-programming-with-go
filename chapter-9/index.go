package main

import "fmt"

func main() {
	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c)

	// message[5] = 'd' string are immutable, cannot assign

	fmt.Println("Printing 'shalom'")
	for i := 0; i < len(message); i++ {
		fmt.Printf("%v\n", message[i])
	}
}