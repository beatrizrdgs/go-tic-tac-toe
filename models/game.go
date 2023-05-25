package models

import "fmt"

type Game struct {
	board  Board
	player [2]Player
}

func NewGame() Game {
	return Game{
		board:  NewBoard(),
		player: [2]Player{Human{symbol: xSymbol}, Computer{symbol: oSymbol}},
	}
}

func (g *Game) PlayGame() {

	b := g.board
	b.resetBoard()
	b.printBoard()

	winner := EMPTY

	for b.checkSpaces() != 0 && winner == EMPTY && !b.checkTie() {
		for _, player := range g.player {
			player.Move(&b)
			winner = b.checkWin()

			if winner != EMPTY || b.checkSpaces() == 0 {
				break
			}
		}
		b.printBoard()
	}
	printWin(winner)
}

func printWin(winner string) {
	if winner != EMPTY {
		fmt.Printf("%s wins\n", winner)
	} else {
		fmt.Println("Tie!")
	}
}
