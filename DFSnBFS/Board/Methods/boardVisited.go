package Methods

var position = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func dfs(board [][]byte, row, col int, visited [][]bool) {
	for _, pos := range position {
		nextRow := row + pos[0]
		nextCol := col + pos[1]

		if nextRow < 0 || nextRow >= len(board) || nextCol < 0 || nextCol >= len(board[0]) {
			continue
		}

		if visited[nextRow][nextCol] || board[nextRow][nextCol] != 'O' {
			continue
		}

		visited[nextRow][nextCol] = true
		dfs(board, nextRow, nextCol, visited)
	}
}

func solve(board [][]byte) {
	rowSize := len(board)
	colSize := len(board[0])
	visited := make([][]bool, rowSize)

	for i := range visited {
		visited[i] = make([]bool, colSize)
	}

	for row := 0; row < rowSize; row++ {
		if board[row][0] == 'O' && !visited[row][0] {
			visited[row][0] = true
			dfs(board, row, 0, visited)
		}
		if board[row][colSize-1] == 'O' && !visited[row][colSize-1] {
			visited[row][colSize-1] = true
			dfs(board, row, colSize-1, visited)
		}
	}

	for col := 1; col < colSize-1; col++ {
		if board[0][col] == 'O' && !visited[0][col] {
			visited[0][col] = true
			dfs(board, 0, col, visited)
		}
		if board[rowSize-1][col] == 'O' && !visited[rowSize-1][col] {
			visited[rowSize-1][col] = true
			dfs(board, rowSize-1, col, visited)
		}
	}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if board[row][col] == 'O' && !visited[row][col] {
				board[row][col] = 'X'
			}
		}
	}
}
