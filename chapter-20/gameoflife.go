package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make(Universe, height)
	for i := 0; i < len(universe); i++ {
		universe[i] = make([]bool, width)
	}
	return universe
}

func (u Universe) Show() {
	for row := 0; row < len(u); row++ {
		for column := 0; column < len(u[row]); column++ {
			cell := u[row][column]
			if cell {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Alive(x, y int) bool {
	wrapY := 0
	wrapX := 0
	universeHeight := height
	universeWidth := width
	if y < 0 {
		wrapY = y + universeHeight
	} else {
		wrapY = y % universeHeight
	}
	if x < 0 {
		wrapX = x + universeWidth
	} else {
		wrapX = x % universeWidth
	}

	return u[wrapY][wrapX]
}

func (u Universe) Neighbors(x, y int) int {
	// get the neighbor coordinates
	neighborsCoords := map[string][]int{
		"topleft":     {y - 1, x - 1},
		"top":         {y - 1, x},
		"topright":    {y - 1, x + 1},
		"left":        {y, x - 1},
		"right":       {y, x + 1},
		"bottomleft":  {y + 1, x - 1},
		"bottom":      {y + 1, x},
		"bottomright": {y + 1, x + 1},
	}

	neighbors := 0

	for _, coords := range neighborsCoords {
		x := coords[1]
		y := coords[0]
		if u.Alive(x, y) {
			neighbors++
		}
	}

	return neighbors
}

func (u Universe) Seed() {
	height := len(u)
	width := len(u[0])
	totalCellsToBeSeeded := (height * width * 25) / 100
	for i := 0; i < totalCellsToBeSeeded; i++ {
		y := rand.Intn(height)
		x := rand.Intn(width)

		if rand.Intn(2) == 1 {
			u[y][x] = true
		} else {
			u[y][x] = false
		}
	}
}

func (u Universe) Next(x, y int) bool {
	cellAlive := u.Alive(x, y)
	neighbors := u.Neighbors(x, y)
	if cellAlive && neighbors < 2 {
		return false
	} else if cellAlive && (neighbors == 2 || neighbors == 3) {
		return true
	} else if cellAlive && neighbors > 3 {
		return false
	} else if !cellAlive && neighbors == 3 {
		return true
	} else {
		return false
	}
}

func Step(a, b Universe) {
	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a[y]); x++ {
			b[y][x] = a.Next(x, y)
		}
	}
	fmt.Print("\033[H")
}

func main() {

	universeA := NewUniverse()
	universeB := NewUniverse()
	universeA.Seed()
	universeA.Show()

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		Step(universeA, universeB)
		universeA.Show()
		universeA, universeB = universeB, universeA
	}
}
