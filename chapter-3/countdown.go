package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10

	for count > 0 { 
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
	
	
	fmt.Println("Unstable launch")
	
	count = 100
	for count > 0 { 
		if rand.Intn(99) == 0 {
			fmt.Println("Launch terminated")
			count = 0
		}

		fmt.Println(count)


		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
}