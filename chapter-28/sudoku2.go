package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

// Grid is a Sudoku grid

type Grid [rows][columns]int8

type SudokuError []error
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit = errors.New("invalid digit")
)

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}


	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column <0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// Error returns one or more errors seperated by commas.
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}



func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de parametres hors limites.")
			break
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}

}