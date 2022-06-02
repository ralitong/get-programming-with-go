package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGophter()
	time.Sleep(4 * time.Second)
}

func sleepyGophter() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}