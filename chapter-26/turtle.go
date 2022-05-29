package main

import (
	"fmt"
	"math/rand"
)

type turtle struct {
	x, y int
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) left() {
	t.x--
}

func (t *turtle) move() {
	for i := 0; i < rand.Intn(100); i++ {
		move := rand.Intn(4)

		switch move {
		case 0:
			t.down()
			break
		case 1:
			t.up()
			break
		case 2:
			t.left()
			break
		case 3:
			t.right()
		default:
			t.down()
		}
	}
}

func (t turtle) printLocation() {
	yDirection := 'N'
	xDirection := 'E'
	if t.y < 0 {
		yDirection = 'S'
	}
	if t.x < 0 {
		xDirection = 'W'
	}

	fmt.Printf("The turtle is %d %c and %d %c\n", t.y, yDirection, t.x, xDirection)
}

func main() {
	oogway := turtle{}
	for i := 0; i < 100; i++ {
		oogway.move()
		oogway.printLocation()
	}

}
