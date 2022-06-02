package main

import (
	"image"
	"log"
	"time"
)

func worker() {
	pos := image.Point{X:10, Y: 10}
	direction := image.Point{X:1, Y:0}
	delay := time.Second
	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			log.Println("current position is",pos)
			delay = delay + time.Second / 2
			next = time.After(delay)
		}

		if pos.X == 30 {
			break
		}
	}
}

func main() {
	worker()
	time.Sleep(30 * time.Second)

}