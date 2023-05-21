package models

import "fmt"

type Human struct {
	symbol string
}

func (h Human) Move(b *Board) {
	validMove := false

	for !validMove {
		var row, column int

		fmt.Print("Enter row and col (e.g: 1 3): ")
		_, err := fmt.Scanf("%d %d\n", &row, &column)
		if err != nil {
			var s string
			fmt.Scan(&s)
			fmt.Println("Invalid move!")
			continue
		}

		row--
		column--

		if column < 0 || column > 2 || row < 0 || row > 2 || b.grid[row][column] != EMPTY {
			fmt.Println("Invalid move!")
			continue
		}

		b.makeMove(NewMove(row, column, h.symbol))
		validMove = true
	}
}
