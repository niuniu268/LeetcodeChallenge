package Methods

func dfs1(board [][]byte, row, col int) {
	for _, pos := range position {
		nextRow := row + pos[0]
		nextCol := col + pos[1]

		if nextRow < 0 || nextRow >= len(board) || nextCol < 0 || nextCol >= len(board[0]) {
			continue
		}

		if board[nextRow][nextCol] != 'O' {
			continue
		}

		board[nextRow][nextCol] = 'A' // Modify to special value
		dfs1(board, nextRow, nextCol)
	}
}

func solveSpecial(board [][]byte) {
	rowSize := len(board)
	colSize := len(board[0])

	// Traverse from left and right sides
	for row := 0; row < rowSize; row++ {
		if board[row][0] == 'O' {
			board[row][0] = 'A'
			dfs1(board, row, 0)
		}
		if board[row][colSize-1] == 'O' {
			board[row][colSize-1] = 'A'
			dfs1(board, row, colSize-1)
		}
	}

	// Traverse from top and bottom sides, we don't need to traverse four corners since they are already traversed in left and right side traversal
	for col := 1; col < colSize-1; col++ {
		if board[0][col] == 'O' {
			board[0][col] = 'A'
			dfs1(board, 0, col)
		}
		if board[rowSize-1][col] == 'O' {
			board[rowSize-1][col] = 'A'
			dfs1(board, rowSize-1, col)
		}
	}

	// Traverse the array, change 'O' to 'X', special value 'A' to 'O'
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == 'A' {
				board[row][col] = 'O'
			}
		}
	}
}
