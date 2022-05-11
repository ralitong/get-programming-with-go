package main

import "fmt"

func putBlackChessPieces(board [8][8]rune) [8][8]rune {
	firstRow := 0
	secondRow := 1

	board[firstRow] = [8]rune{'r', 'h', 'b', 'k', 'q', 'b', 'h', 'r'}
	board[secondRow] = [8]rune{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'}
	return board
}

func putWhiteChessPieces(board [8][8]rune) [8][8]rune {
	lastRow := len(board) - 1
	secondLastRow := len(board) - 2

	board[secondLastRow] = [8]rune{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'}
	board[lastRow] = [8]rune{'R', 'H', 'B', 'K', 'Q', 'B', 'H', 'R'}
	return board
}

func displayBoard(board [8][8]rune) {
	for column := 0; column < len(board[0]); column++ {
		fmt.Print("____")
	}
	fmt.Println()
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board[0]); column++ {
			if board[row][column] != 0 {
				fmt.Printf("| %c ", board[row][column])
			} else {
				fmt.Printf("|   ")
			}
		}
		fmt.Print("|")
		fmt.Println()
		for column := 0; column < len(board[0]); column++ {
			fmt.Print("____")
		}
		fmt.Println()
	}
}

func main() {
	var board [8][8]rune
	halfBoard := putBlackChessPieces(board)
	fullBoard := putWhiteChessPieces(halfBoard)
	displayBoard(fullBoard)

}
