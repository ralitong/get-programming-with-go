package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	speed = int(50)
)

type Earth struct {
	receiver chan string
}

func (e *Earth) receiveMessage(message string) {
	e.receiver <- message
}

func (e *Earth) printMessages() {
	updateInterval := 250 * time.Millisecond
	nextCycle := time.After(updateInterval)
	for {
		select {
		case e := <-e.receiver:
			log.Printf("Earth receives message: %v", e)
		case <- nextCycle:
			log.Print("Earth not receiving any message")
			nextCycle = time.After(updateInterval)
		}
	}
}

// Roverdirver drives a rover around the surface of Mars.
type RoverDriver struct {
	name string
	mars *MarsGrid
	earth *Earth
	pos image.Point
}

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine
func (r *RoverDriver) drive() {
	direction := image.Point{X: 1, Y: 0}
	currCommand := right
	pos := r.pos.Add(direction)
	for {

		if r.move(pos) {
			// log.Printf("%s moved to %v", r.name,r.pos)
			pos = pos.Add(direction)
			
		} else {
			
			currCommand = command(rand.Intn(1)+1)

			switch currCommand {
			case right:
				direction = image.Point{X: -direction.Y, Y: direction.X}
			case left:
				direction = image.Point{X: direction.Y, Y: -direction.X}
			}
			// log.Printf("%s could not move to position %v, direction %v",r.name, r.pos, direction)
			pos = pos.Add(direction)
		}

		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func (r *RoverDriver) move(p image.Point) bool {
	if r.mars.occupy(r,p) {
		r.pos = p
		return true
	}
	return false
}

func (r *RoverDriver) land() {
	sizeOfMars := len(r.mars.grid)
	randomPoint := image.Point{X: rand.Intn(sizeOfMars), Y: rand.Intn(sizeOfMars)}
	landed := false

	for !landed {
		if r.mars.occupy(r, randomPoint) {
			r.pos = randomPoint
			landed = true
		}
	}
}

func(r *RoverDriver) transmitLifeToEarth() {
	messages := make([]string, 0)
	for {
		life := r.mars.getLife(r.pos)
		if (life > 900) {
			messages = append(messages, fmt.Sprintf("Rover %s found life value %d on %v", r.name, life, r.pos))
		}

		if len(messages) == 5 {
			for _, message := range messages {
				r.earth.receiveMessage(message)
			}
			messages = nil
		}

		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func NewRoverDriver(name string, g *MarsGrid, e *Earth) *RoverDriver {
	r := &RoverDriver{
		name: name,
		mars: g,
		earth: e,
	}
	r.land()
	return r
}

type Cell struct {
	rover *RoverDriver
	life int
}

func (c *Cell) hasRover() bool {
	return c.rover != nil
}

// MarsGrid represents a grid of some of the surface
// of Mars. It may be used concurrently by different
// goroutines.
type MarsGrid struct {
	grid [][]Cell
	occupymu sync.Mutex
}

func (g *MarsGrid) lock() {
	g.occupymu.Lock()
}

func (g *MarsGrid) unlock() {
	g.occupymu.Unlock()
}

// Occupy occupies a cell at the given point in the grid. It
// returns nil if the point is already occupied or the point is
// outside the grid. Otherwise it returns a value that can be
// used to move to different places on the grid.
func (g *MarsGrid) occupy(rover *RoverDriver,p image.Point) bool{
	g.lock()
	defer g.unlock()

	if p.X < 0 || p.X >= len(g.grid) || p.Y < 0 || p.Y >= len(g.grid) {
		return false
	} 

	if(g.grid[p.Y][p.X].rover == nil) {
		g.grid[p.Y][p.X].rover = rover
		g.grid[rover.pos.Y][rover.pos.X].rover = nil
		return true
	} else {
		return false
	}
}

func (g *MarsGrid) display() {
	gridSize := len(g.grid)

	for {
		for row := 0; row < gridSize; row++ {
			for column := 0; column < gridSize; column++ {
				cell := g.grid[row][column]
				if cell.hasRover() {
					fmt.Printf("%s",cell.rover.name[0:1])
					// log.Printf("cell x:%d y:%d has rover", column, row)
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	
		time.Sleep(time.Duration(speed) * time.Millisecond)
		fmt.Print("\033[H")
	}
}

func (g *MarsGrid) getLife(p image.Point) int{
	return g.grid[p.Y][p.X].life
}

func NewMarsGrid() MarsGrid {
	worldSize := 100
	grid := make([][]Cell, worldSize)
	for i := 0; i < worldSize; i++ {
		grid[i] = make([]Cell, worldSize)
	}
	for i := 0; i < worldSize; i++ {
		for j:= 0; j < worldSize; j++ {
			grid[i][j] = Cell{life: rand.Intn(1000)+1}
		}
	}

	return MarsGrid{grid: grid}
}


func main() {
	mars := NewMarsGrid()
	earth := Earth{receiver: make(chan string)}
	for _, i := range "CURIOSITY"{
		rover := NewRoverDriver(string(i), &mars, &earth)
		go rover.drive()
		go rover.transmitLifeToEarth()
	}
	// go mars.display()
	go earth.printMessages()
	time.Sleep(100 * time.Second)
}
