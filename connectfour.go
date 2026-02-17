package connectfour

import (
	"fmt"
)

const rows, columns = 6, 7

var Board [columns][rows]int

type Move struct {
	Player int
	Column int
	Row    int
}

type Direction struct {
	Vertical   int
	Horizontal int
}

func updateBoard(column, player int) Move {
	row := 0
	// Find first empty square in column
	for Board[column][row] != 0 {
		row++
	}
	Board[column][row] = player
	return Move{player, column, row}
}

func printBoard() {
	for r := range rows {
		currentRow := rows - 1 - r
		for c := range columns {
			currentColumn := c

			var char string
			switch Board[currentColumn][currentRow] {
			case 0:
				char = " "
			case 1:
				char = "x"
			case 2:
				char = "o"
			}

			fmt.Print("|", char)
		}
		fmt.Println("|")
	}
	fmt.Println("--------------------")
}

func isWon(lastMove Move) bool {
	horizontal := Direction{0, 1}
	vertical := Direction{1, 0}
	diagonal := Direction{1, 1}

	check1 := isWonDirection(lastMove, horizontal)
	check2 := isWonDirection(lastMove, vertical)
	check3 := isWonDirection(lastMove, diagonal)

	return check1 || check2 || check3
}

func isWonDirection(lastMove Move, direction Direction) bool {
	length := 1

	currentColumn := lastMove.Column + direction.Horizontal
	currentRow := lastMove.Row + direction.Vertical

	for isInbounds(currentColumn, currentRow) && (Board[currentColumn][currentRow] == lastMove.Player) {
		length++
		currentColumn = currentColumn + direction.Horizontal
		currentRow = currentRow + direction.Vertical
	}

	currentColumn = lastMove.Column - direction.Horizontal
	currentRow = lastMove.Row - direction.Vertical

	for isInbounds(currentColumn, currentRow) && (Board[currentColumn][currentRow] == lastMove.Player) {
		length++
		currentColumn = currentColumn - direction.Horizontal
		currentRow = currentRow - direction.Vertical
	}
	return length > 3
}

func isInbounds(column, row int) bool {
	return (column < columns) && (column >= 0) && (row < rows) && (row >= 0)
}

func main() {
	player := 1
	var isOver bool
	var turns int

	var column int

	printBoard()

	for !isOver {
		fmt.Printf("Player %d turn plese choose a column (0-6)\n", player)

		_, err := fmt.Scan(&column)
		if err != nil {
			fmt.Println("Invalid input please enter a number")
			continue
		} else if column >= columns || column < 0 {
			fmt.Println("Invalid input please enter a number 0-6")
			continue
		}

		last_move := updateBoard(column, player)
		turns++
		printBoard()

		if isWon(last_move) {
			fmt.Printf("Congrats Player %d you won\n", player)
			isOver = true
		} else if turns == (rows * columns) {
			fmt.Println("Game is drawn ")
			isOver = true
		}
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}

}
