package main

import (
	"errors"
	"fmt"
)


type Sudoku struct {
	grid      [][]int
	fixedGrid [][]int
}

func (s Sudoku) inBounds(x, y int) bool {
	height := len(s.grid)
	width := len(s.grid)
	if y < 0 || y >= height {
		return false
	}
	if x < 0 || y >= width {
		return false
	}
	return true
}

func (s *Sudoku) backupGrid() {
	height := len(s.grid)
	width := len(s.grid)
	fixedGrid := make([][]int, height)
	for y := 0; y < height; y++ {
		fixedGrid[y] = make([]int, width)
		for x := 0; x < width; x++ {
			fixedGrid[y][x] = s.grid[y][x]
		}
	}

	s.fixedGrid = fixedGrid
}

func (s Sudoku) canBeCleared(x, y int) bool {
	return s.fixedGrid[y][x] == 0
}

func (s Sudoku) doesValueAppearHorizontally(x, y, val int) bool {
	row := s.grid[y]
	for i := 0; i < len(row); i++ {
		if i == x {
			continue
		}

		if row[i] == val {
			return true
		}
	}

	return false
}

func (s Sudoku) doesValueAppearVertically(x, y, val int) bool {
	height := len(s.grid)

	for column := 0; column < height; column++ {
		if column == y {
			continue
		}
		if s.grid[column][x] == val {
			return true
		}
	}
	return false
}

func (s Sudoku) doesValueAppearIn3x3Grid(x, y, val int) bool {
	xOrigin := s.getNearestOrigin(x)
	yOrigin := s.getNearestOrigin(y)

	for row := yOrigin; row < yOrigin+3; row++ {
		for column := xOrigin; column < xOrigin+3; column++ {
			if row == y && column == x {
				continue
			}

			if s.grid[row][column] == val {
				return true
			}
		}
	}

	return false
}

func (s Sudoku) getNearestOrigin(num int) int {
	limit := len(s.grid)
	nearestOrigin := 0

	for i := 0; i < limit; i += 3 {

		distance := num - i
		if distance < 0 {
			distance *= -1
		}

		if distance <= 2 {
			nearestOrigin = i
			break
		}
	}

	return nearestOrigin
}

func (s Sudoku) isNumberBetweenOneAndNine(num int) bool {
	return num >= 1 && num <= 9
}

func (s *Sudoku) put(x, y, val int) error {
	if !s.isNumberBetweenOneAndNine(val) {
		return errors.New(fmt.Sprintf("%d is not between 1 and 9", val))
	}

	if !s.inBounds(x, y) {
		return errors.New(fmt.Sprintf("x:%d and y:%d is out of bounds", x, y))
	}

	if s.doesValueAppearHorizontally(x, y, val) {
		return errors.New(fmt.Sprintf("%d appears horizontally", val))
	}

	if s.doesValueAppearVertically(x, y, val) {
		return errors.New(fmt.Sprintf("%d appears vertically", val))
	}

	if s.doesValueAppearIn3x3Grid(x, y, val) {
		return errors.New(fmt.Sprintf("%d appears in 3x3 grid", val))
	}

	if !s.canBeCleared(x, y) {
		return errors.New(fmt.Sprintf("x:%d and y:%d already has a fixed value: %d",
			x,
			y,
			s.grid[y][x]))
	}

	s.grid[y][x] = val

	return nil
}

func NewSudoku(grid [][]int) Sudoku {
	s := Sudoku{grid, nil}
	s.backupGrid()
	return s
}

func putValuesToSudoku(s Sudoku, x,y,val int) {
	err := s.put(x,y,val)

	if err != nil {
		fmt.Printf("The error is: %v\n",err)
	}
}

func main() {
	s := NewSudoku([][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
		})


		putValuesToSudoku(s, 10, 10, 5)
		putValuesToSudoku(s, 3, 0, 12)
		putValuesToSudoku(s, 8, 0, 7)
		putValuesToSudoku(s, 8, 0, 6)
		putValuesToSudoku(s, 2, 0, 9)
		putValuesToSudoku(s, 0, 0, 5)



}
