package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGophter(i,c)
	}

	for i := 0; i < 5; i++ {
		gopherID := <- c
		fmt.Println("gopher",gopherID,"has finished sleeping")
	}
	
}


func sleepyGophter(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...",id,"snore ...")
	c <- id
}