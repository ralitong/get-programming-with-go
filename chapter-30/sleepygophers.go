package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go sleepyGophter(i)
	}

	time.Sleep(4 * time.Second)
}

func sleepyGophter(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...",id,"snore ...")
}