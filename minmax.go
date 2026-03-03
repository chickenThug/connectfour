package main

func isWon(board *[columns][rows]int, lastMove Move) bool {
	horizontal := Direction{1, 0}
	vertical := Direction{0, 1}
	diagonal := Direction{1, 1}
	antiDiagonal := Direction{1, -1}

	check1 := isWonDirection(board, lastMove, horizontal)
	check2 := isWonDirection(board, lastMove, vertical)
	check3 := isWonDirection(board, lastMove, diagonal)
	check4 := isWonDirection(board, lastMove, antiDiagonal)

	return check1 || check2 || check3 || check4
}

func isWonDirection(board *[columns][rows]int, lastMove Move, direction Direction) bool {
	length := 1

	currentColumn := lastMove.Column + direction.Horizontal
	currentRow := lastMove.Row + direction.Vertical

	for isInbounds(currentColumn, currentRow) && (board[currentColumn][currentRow] == lastMove.Player) {
		length++
		currentColumn = currentColumn + direction.Horizontal
		currentRow = currentRow + direction.Vertical
	}

	currentColumn = lastMove.Column - direction.Horizontal
	currentRow = lastMove.Row - direction.Vertical

	for isInbounds(currentColumn, currentRow) && (board[currentColumn][currentRow] == lastMove.Player) {
		length++
		currentColumn = currentColumn - direction.Horizontal
		currentRow = currentRow - direction.Vertical
	}
	return length > 3
}

// heuristic value of the board state
func value(board *[columns][rows]int) int {
	horizontal := Direction{1, 0}
	vertical := Direction{0, 1}
	diagonal := Direction{1, 1}
	antiDiagonal := Direction{1, -1}

	directions := [4]Direction{horizontal, vertical, diagonal, antiDiagonal}
	value := 0

	for row := range rows {
		for column := range columns {
			for _, dir := range directions {
				endColumn := column + (winLength-1)*dir.Horizontal
				endRow := row + (winLength-1)*dir.Vertical

				// skip if out of bounds
				if endColumn < 0 || endRow < 0 || endColumn >= columns || endRow >= rows {
					continue
				}
				value += windowValue(board, column, row, dir)
			}
		}
	}

	return value
}

func windowValue(board *[columns][rows]int, column, row int, dir Direction) int {
	player1, player2 := 0, 0
	for i := range winLength {
		if i != 0 {
			column += dir.Horizontal
			row += dir.Vertical
		}
		marker := board[column][row]
		switch marker {
		case -1:
			player2 += 1
		case 1:
			player1 += 1
		}
	}
	if player1 > 0 && player2 > 0 {
		return 0
	} else if player1 > 0 {
		return player1
	} else if player2 > 0 {
		return -player2
	} else {
		return 0
	}
}

func isTerminal(board *[columns][rows]int, lastMove Move) int {
	if lastMove.Player == 0 {
		return 0
	}
	if isWon(board, lastMove) {
		return 2
	}
	for column := range columns {
		if board[column][rows-1] == 0 {
			return 0
		}
	}
	return 1
}

func getMoves(board *[columns][rows]int, player int) []Move {
	var moves []Move
	for column := range columns {
		for row := range rows {
			if board[column][row] == 0 {
				moves = append(moves, Move{player, column, row})
				break
			}
		}
	}
	return moves
}

func minimax(board *[columns][rows]int, lastMove Move, depth int, player int) int {
	isOver := isTerminal(board, lastMove)
	if isOver > 0 {
		if isOver == 2 {
			return 999999999999999 * lastMove.Player
		} else {
			return 0
		}
	}
	if depth == 0 {
		return value(board)
	}
	moves := getMoves(board, player)
	if player == 1 {
		value := -9999999999999999
		for _, move := range moves {
			board[move.Column][move.Row] = player
			value = max(value, minimax(board, move, depth-1, player*-1))
			board[move.Column][move.Row] = 0
		}
		return value
	} else {
		value := 9999999999999999
		for _, move := range moves {
			board[move.Column][move.Row] = player
			value = min(value, minimax(board, move, depth-1, player*-1))
			board[move.Column][move.Row] = 0
		}
		return value
	}
}

func GetBotMove(board *[columns][rows]int, depth int, player int) int {
	moves := getMoves(board, player)
	bestColumn := 0
	var bestValue int
	if player == 1 {
		bestValue = -9999999999999999
		for _, move := range moves {
			board[move.Column][move.Row] = player
			value := minimax(board, move, depth, player*-1)
			board[move.Column][move.Row] = 0

			if value > bestValue {
				bestColumn = move.Column
				bestValue = value
			}
		}
	} else {
		bestValue = 9999999999999999
		for _, move := range moves {
			board[move.Column][move.Row] = player
			value := minimax(board, move, depth, player*-1)
			board[move.Column][move.Row] = 0

			if value < bestValue {
				bestColumn = move.Column
				bestValue = value
			}
		}
	}
	return bestColumn
}
