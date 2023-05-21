package models

type Move struct {
	row    int
	column int
	symbol string
}

func (b *Board) makeMove(m Move) {
	b.grid[m.row][m.column] = m.symbol
}

func NewMove(row int, column int, symbol string) Move {
	return Move{
		row:    row,
		column: column,
		symbol: symbol,
	}
}

func NewEmptyMove(row int, column int) Move {
	return Move{
		row:    row,
		column: column,
		symbol: EMPTY,
	}
}
