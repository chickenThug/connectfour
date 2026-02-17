package main

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

func update_board(column, player int) Move {
	row := 0
	// Find first empty square in column
	for Board[column][row] != 0 {
		row++
	}
	Board[column][row] = player
	return Move{player, column, row}
}

func print_board() {
	for r := range rows {
		current_row := rows - 1 - r
		for c := range columns {
			current_column := c

			var char string
			switch Board[current_column][current_row] {
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

func is_won(last_move Move) bool {
	horizontal := Direction{0, 1}
	vertical := Direction{1, 0}
	diagonal := Direction{1, 1}

	check1 := is_won_direction(last_move, horizontal)
	check2 := is_won_direction(last_move, vertical)
	check3 := is_won_direction(last_move, diagonal)

	return check1 || check2 || check3
}

func is_won_direction(last_move Move, direction Direction) bool {
	length := 1

	current_column := last_move.Column + direction.Horizontal
	current_row := last_move.Row + direction.Vertical

	for is_inbounds(current_column, current_row) && (Board[current_column][current_row] == last_move.Player) {
		length++
		current_column = current_column + direction.Horizontal
		current_row = current_row + direction.Vertical
	}

	current_column = last_move.Column - direction.Horizontal
	current_row = last_move.Row - direction.Vertical

	for is_inbounds(current_column, current_row) && (Board[current_column][current_row] == last_move.Player) {
		length++
		current_column = current_column - direction.Horizontal
		current_row = current_row - direction.Vertical
	}
	return length > 3
}

func is_inbounds(column, row int) bool {
	return (column < columns) && (column >= 0) && (row < rows) && (row >= 0)
}

func main() {
	player := 1
	var is_over bool
	var turns int

	var column int

	print_board()

	for !is_over {
		fmt.Printf("Player %d turn plese choose a column (0-6)\n", player)

		_, err := fmt.Scan(&column)
		if err != nil {
			fmt.Println("Invalid input please enter a number")
			continue
		} else if column >= columns || column < 0 {
			fmt.Println("Invalid input please enter a number 0-6")
			continue
		}

		last_move := update_board(column, player)
		turns++
		print_board()

		if is_won(last_move) {
			fmt.Printf("Congrats Player %d you won\n", player)
			is_over = true
		} else if turns == (rows * columns) {
			fmt.Println("Game is drawn ")
			is_over = true
		}
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}

}
