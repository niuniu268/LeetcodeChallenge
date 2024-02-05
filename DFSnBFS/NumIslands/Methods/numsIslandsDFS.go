package Methods

func numIslands(grid [][]byte) int {
	// 用1标记已访问
	visited := make([][]int, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	var bfs func(x, y int)
	bfs = func(x, y int) {
		stack := make([][]int, 0)
		stack = append(stack, []int{x, y})
		moveX := []int{1, -1, 0, 0}
		moveY := []int{0, 0, 1, -1}

		for len(stack) != 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for i := 0; i < 4; i++ {
				dx := moveX[i] + node[0]
				dy := moveY[i] + node[1]
				if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[0]) || visited[dx][dy] == 1 {
					continue
				}
				visited[dx][dy] = 1
				if grid[dx][dy] == '1' {
					stack = append(stack, []int{dx, dy})
				}
			}
		}
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if visited[i][j] == 0 && grid[i][j] == '1' {
				bfs(i, j)
				visited[i][j] = 1
				result++
			}
		}
	}

	return result
}
