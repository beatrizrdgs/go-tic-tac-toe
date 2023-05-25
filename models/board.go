package models

import "fmt"

type Board struct {
	grid [3][3]string
}

var PossibleWins = [8][3][2]int{
	// rows
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	// columns
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	// diagonals
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
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

	players := []string{"X", "O"}

	for _, player := range players {
		for _, win := range PossibleWins {
			winner := true
			for _, position := range win {
				row, column := position[0], position[1]
				if b.grid[row][column] != player {
					winner = false
					break
				}
			}
			if winner {
				return player
			}
		}
	}
	return EMPTY
}

func (b *Board) checkTie() bool {
	isTie := true
	for _, line := range PossibleWins {
		symbol1 := b.grid[line[0][0]][line[0][1]]
		winIsPossible := true
		for _, position := range line[1:] {
			symbol2 := b.grid[position[0]][position[1]]
			if symbol1 != symbol2 && symbol1 != EMPTY && symbol2 != EMPTY {
				winIsPossible = false
				break
			}
			if symbol2 != EMPTY {
				symbol1 = symbol2
			}
		}
		if winIsPossible {
			isTie = false
			break
		}
	}
	return isTie
}
