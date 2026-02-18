package main

import "fmt"

type Move struct {
	Player int
	Column int
	Row    int
}

type Direction struct {
	Vertical   int
	Horizontal int
}

type Game struct {
	Player int
	Turns  int
	IsOver bool
	Board  [columns][rows]int
}

func NewGame() Game {
	return Game{Player: 1}
}

func (g *Game) updateBoard(column int) Move {
	row := 0
	// Find first empty square in column
	for g.Board[column][row] != 0 {
		row++
	}
	g.Board[column][row] = g.Player
	return Move{g.Player, column, row}
}

func (g *Game) printBoard() {
	for r := range rows {
		currentRow := rows - 1 - r
		for c := range columns {
			currentColumn := c

			var char string
			switch g.Board[currentColumn][currentRow] {
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

func (g *Game) isWon(lastMove Move) bool {
	horizontal := Direction{0, 1}
	vertical := Direction{1, 0}
	diagonal := Direction{1, 1}
	antiDiagonal := Direction{1, -1}

	check1 := g.isWonDirection(lastMove, horizontal)
	check2 := g.isWonDirection(lastMove, vertical)
	check3 := g.isWonDirection(lastMove, diagonal)
	check4 := g.isWonDirection(lastMove, antiDiagonal)

	return check1 || check2 || check3 || check4
}

func (g *Game) isWonDirection(lastMove Move, direction Direction) bool {
	length := 1

	currentColumn := lastMove.Column + direction.Horizontal
	currentRow := lastMove.Row + direction.Vertical

	for isInbounds(currentColumn, currentRow) && (g.Board[currentColumn][currentRow] == lastMove.Player) {
		length++
		currentColumn = currentColumn + direction.Horizontal
		currentRow = currentRow + direction.Vertical
	}

	currentColumn = lastMove.Column - direction.Horizontal
	currentRow = lastMove.Row - direction.Vertical

	for isInbounds(currentColumn, currentRow) && (g.Board[currentColumn][currentRow] == lastMove.Player) {
		length++
		currentColumn = currentColumn - direction.Horizontal
		currentRow = currentRow - direction.Vertical
	}
	return length > 3
}

func isInbounds(column, row int) bool {
	return (column < columns) && (column >= 0) && (row < rows) && (row >= 0)
}

func (g *Game) isColumnFull(column int) bool {
	return g.Board[column][rows-1] != 0
}

func (g *Game) run() {
	var column int

	g.printBoard()

	for !g.IsOver {
		fmt.Printf("Player %d turn plese choose a column (0-6)\n", g.Player)

		_, err := fmt.Scan(&column)
		if err != nil {
			fmt.Println("Invalid input please enter a number")
			continue
		} else if column >= columns || column < 0 {
			fmt.Println("Invalid input please enter a number 0-6")
			continue
		} else if g.isColumnFull(column) {
			fmt.Println("The choosen column is full, please choose another")
		}

		last_move := g.updateBoard(column)
		g.Turns++
		g.printBoard()

		if g.isWon(last_move) {
			fmt.Printf("Congrats Player %d you won\n", g.Player)
			g.IsOver = true
		} else if g.Turns == (rows * columns) {
			fmt.Println("Game is drawn ")
			g.IsOver = true
		}
		if g.Player == 1 {
			g.Player = 2
		} else {
			g.Player = 1
		}
	}
}
