package main

import (
	"fmt"
	"math/rand"
)

const (
	PLAYER = "X"
	AI     = "O"
)

var (
	board [3][3]string
)

func main() {

	playGame()

}

func resetBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Printf("%s | %s | %s\n", board[0][0], board[0][1], board[0][2])
	fmt.Printf("——|———|——\n")
	fmt.Printf("%s | %s | %s\n", board[1][0], board[1][1], board[1][2])
	fmt.Printf("——|———|——\n")
	fmt.Printf("%s | %s | %s\n", board[2][0], board[2][1], board[2][2])
}

func checkSpaces() int {
	var spaces = 9
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != " " {
				spaces--
			}
		}
	}
	return spaces
}

func playerMove() {
	var validMove = false

	for !validMove {
		var row, col int

		fmt.Print("Enter row and col (e.g: 1 3): ")
		n, err := fmt.Scanf("%d %d\n", &row, &col)
		if err != nil || n != 2 {
			var s string
			fmt.Scanf("%s", &s)
			fmt.Println("Invalid move!")
			continue
		}

		row--
		col--

		if col < 0 || col > 2 || row < 0 || row > 2 || board[row][col] != " " {
			var s string
			fmt.Scanf("%s", &s)
			fmt.Println("Invalid move!")
			continue
		}

		board[row][col] = PLAYER
		validMove = true
	}
}

func checkWin() string {
	// Row
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
	}
	// Column
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return board[0][i]
		}
	}
	// Diagonal
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}
	return " "
}

func aiMove() {
	// Win
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				board[i][j] = AI
				if checkWin() == AI {
					return
				}
				board[i][j] = " "
			}
		}
	}

	// Block
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				board[i][j] = PLAYER
				if checkWin() == PLAYER {
					board[i][j] = AI
					return
				}
				board[i][j] = " "
			}
		}
	}

	// Random
	var row, col int
	for {
		row = rand.Intn(3)
		col = rand.Intn(3)
		if board[row][col] == " " {
			break
		}
	}
	board[row][col] = AI
}

func printWin(winner string) {
	if winner == PLAYER {
		fmt.Println("You win!")
	} else if winner == AI {
		fmt.Println("You lose!")
	} else {
		fmt.Println("Tie!")
	}
}

func playGame() {

	resetBoard()
	var winner = " "

	for checkSpaces() != 0 && winner == " " {
		printBoard()

		playerMove()
		winner = checkWin()

		if winner != " " || checkSpaces() == 0 {
			break
		}

		aiMove()
		winner = checkWin()

		if winner != " " || checkSpaces() == 0 {
			break
		}
	}

	printBoard()
	printWin(winner)
}
