package models

import (
	"math/rand"
	"time"
)

type Computer struct {
	symbol string
}

func (c Computer) Move(b *Board) {
	// Win
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if b.grid[row][column] == EMPTY {
				b.makeMove(NewMove(row, column, c.symbol))
				if b.checkWin() == oSymbol {
					return
				}
				b.makeMove(NewEmptyMove(row, column))
			}
		}
	}

	// Block
	oppositeSymbol := xSymbol
	if oppositeSymbol == c.symbol {
		oppositeSymbol = oSymbol
	}

	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if b.grid[row][column] == EMPTY {
				b.makeMove(NewMove(row, column, oppositeSymbol))
				if b.checkWin() == xSymbol {
					b.makeMove(NewMove(row, column, c.symbol))
					return
				}
				b.makeMove(NewEmptyMove(row, column))
			}
		}
	}

	// Random
	rand.Seed(time.Now().UnixNano())
	var row, column int
	for {
		row = rand.Intn(3)
		column = rand.Intn(3)
		if b.grid[row][column] == EMPTY {
			break
		}
	}
	b.makeMove(NewMove(row, column, c.symbol))
}
