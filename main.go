package main

import (
	"fmt"
)

const rows, columns = 6, 7

var Board [columns][rows]int

func update_board(column, player int) {
	i := 0
	// Find first empty square in column
	for Board[column][i] != 0 {
		i++
	}
	Board[column][i] = player
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

func main() {
	print_board()
	update_board(3, 1)
	print_board()
	update_board(3, 2)
	print_board()
}
