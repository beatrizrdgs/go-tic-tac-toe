package models

import "fmt"

type Board struct {
	grid [3][3]string
}

func NewBoard() Board {
	b := Board{}
	b.resetBoard()
	return b
}

func (b *Board) resetBoard() {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			b.makeMove(NewEmptyMove(row, column))
		}
	}
}

func (b *Board) printBoard() {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			fmt.Printf("%s ", b.grid[row][column])
			if column < 2 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if row < 2 {
			fmt.Println("——|———|——")
		}
	}
}

func (b *Board) checkSpaces() int {
	var spaces = 9
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if b.grid[row][column] != EMPTY {
				spaces--
			}
		}
	}
	return spaces
}

func (b *Board) checkWin() string {

	for row := 0; row < 3; row++ {
		if b.grid[row][0] == b.grid[row][1] && b.grid[row][0] == b.grid[row][2] {
			return b.grid[row][0]
		}
	}

	for column := 0; column < 3; column++ {
		if b.grid[0][column] == b.grid[1][column] && b.grid[0][column] == b.grid[2][column] {
			return b.grid[0][column]
		}
	}
	// Diagonal
	if b.grid[0][0] == b.grid[1][1] && b.grid[0][0] == b.grid[2][2] {
		return b.grid[0][0]
	}
	if b.grid[0][2] == b.grid[1][1] && b.grid[0][2] == b.grid[2][0] {
		return b.grid[0][2]
	}
	return EMPTY
}
