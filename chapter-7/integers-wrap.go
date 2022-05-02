package main

import "fmt"

func main() {
	const red uint8 = 255
	const number int8 = 127
	
	tempRed := red
	tempNumber := number
	
	fmt.Println(tempRed + 2)
	fmt.Println(tempNumber + 3)


	tempRed = 0
	tempNumber = -128
	fmt.Println(tempRed-1)
	fmt.Println(tempNumber-1)

	var max uint16 = 65535
	fmt.Println(max+1)
}