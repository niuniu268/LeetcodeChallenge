package Methods

var position = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func dfs(heights [][]int, row, col, sign int, visited [][][]bool) {
	for _, current := range position {
		curRow := row + current[0]
		curCol := col + current[1]

		if curRow < 0 || curRow >= len(heights) || curCol < 0 || curCol >= len(heights[0]) {
			continue
		}

		if heights[curRow][curCol] < heights[row][col] || visited[curRow][curCol][sign] {
			continue
		}

		visited[curRow][curCol][sign] = true
		dfs(heights, curRow, curCol, sign, visited)
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	rowSize := len(heights)
	colSize := len(heights[0])
	var ans [][]int

	visited := make([][][]bool, rowSize)
	for i := range visited {
		visited[i] = make([][]bool, colSize)
		for j := range visited[i] {
			visited[i][j] = make([]bool, 2)
		}
	}

	for row := 0; row < rowSize; row++ {
		visited[row][colSize-1][0] = true
		visited[row][0][1] = true
		dfs(heights, row, colSize-1, 0, visited)
		dfs(heights, row, 0, 1, visited)
	}

	for col := 0; col < colSize; col++ {
		visited[rowSize-1][col][0] = true
		visited[0][col][1] = true
		dfs(heights, rowSize-1, col, 0, visited)
		dfs(heights, 0, col, 1, visited)
	}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if visited[row][col][0] && visited[row][col][1] {
				ans = append(ans, []int{row, col})
			}
		}
	}
	return ans
}
