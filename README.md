# Tic-Tac-Toe
This is a Command Line Interface Tic Tac Toe game implemented in Go using structs, interfaces and pointers, where the player can play against an artificial intelligence. 

## Installation & Usage
To run this program, you need to have the Go compiler installed on your system. You can follow the installation instructions on [Go's website](https://go.dev/)

To get started with the program, you need to first clone or download the repository. You can do this by clicking on the "Code" button at the top of the repository page and selecting "Download ZIP" to download a compressed file of the repository, or copy the repository URL and run the following command in your terminal:
```
git clone https://github.com/beatrizhub/go-tic-tac-toe
```

Once you have the compiler installed and cloned this repository, move into the program's directory and run it using one of the following command:

```
go run .
```

## Gameplay
When the game starts, enter the row and column (1, 2, or 3) of the cell where you want to place your marker separated by a single space (e.g. 1 3). After your move, the AI will make its move and the game board will be displayed.

The game continues until there's a tie or either the player or the AI wins.

## Algorithm
The AI in this program uses uses a simple strategy to select its moves. It first checks for any winning move by iterating through all the empty cells on the board and placing the computer's symbol in each one. If a move results in a win, the function immediately returns and the move is made.

If there is no winning move available, the function checks for a blocking move. It iterates through all the empty cells on the board again and places the player's symbol in each one. If a move results in blocking the player from winning, the computer places its symbol in that cell and returns, making the move.

If neither a winning nor a blocking move is available, the function makes a random move. It generates a random row and column until an empty cell is found, and then places the computer's symbol in that cell.

Overall, this function prioritizes winning over blocking the player's moves, and random moves are made only when no winning or blocking moves are available.

## License
This program is licensed under the MIT License. Feel free to use and modify this program for your own purposes.