package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop = command(3)
)

// Roverdirver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc chan command
	stopped bool
}

// Left turns the rover left (90 deg clockwise)
func (r *RoverDriver) left() {
	r.commandc <- left
}

// Right turns the rover right (90 deg clockwise)
func (r *RoverDriver) right() {
	r.commandc <- right
}

func (r *RoverDriver) start() {
	r.commandc <- start
}

func (r * RoverDriver) stop() {
	r.commandc <- stop
}

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{X: -direction.Y, Y: direction.X}
			case left:
				direction = image.Point{X: direction.Y, Y: -direction.X}
			case stop:
				r.stopped = true
			case start:
				r.stopped = false
			}
			log.Printf("new direction %v", direction)
		case <- nextMove:
			if r.stopped {
				log.Println("Rover is not moving ...")
			} else {
				pos = pos.Add(direction)
				log.Printf("moved to %v", pos)
			}
			nextMove = time.After(updateInterval)
		}
	}
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		stopped: false,
	}
	go r.drive()
	return r
}


func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.left()
	time.Sleep(3 * time.Second)
	r.right()
	time.Sleep(3 * time.Second)
	r.stop()
	time.Sleep(3 * time.Second)
	r.start()
	time.Sleep(3 * time.Second)

}
