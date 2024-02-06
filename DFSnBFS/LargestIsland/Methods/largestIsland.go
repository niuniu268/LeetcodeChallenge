package Methods

var dir = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} // 四个方向

func dfs(grid [][]int, visited [][]bool, x, y, mark int) int {
	if visited[x][y] || grid[x][y] == 0 {
		return 0
	}

	visited[x][y] = true
	grid[x][y] = mark
	count := 1

	for _, d := range dir {
		nextX := x + d[0]
		nextY := y + d[1]

		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}

		count += dfs(grid, visited, nextX, nextY, mark)
	}

	return count
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	gridNum := make(map[int]int)
	mark := 2
	isAllGrid := true

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				isAllGrid = false
			}
			if !visited[i][j] && grid[i][j] == 1 {
				count := dfs(grid, visited, i, j, mark)
				gridNum[mark] = count
				mark++
			}
		}
	}

	if isAllGrid {
		return n * m
	}

	result := 0
	visitedGrid := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := 1
			visitedGrid = make(map[int]bool)

			if grid[i][j] == 0 {
				for _, d := range dir {
					nearI := i + d[0]
					nearJ := j + d[1]

					if nearI < 0 || nearI >= n || nearJ < 0 || nearJ >= m {
						continue
					}

					if visitedGrid[grid[nearI][nearJ]] {
						continue
					}

					count += gridNum[grid[nearI][nearJ]]
					visitedGrid[grid[nearI][nearJ]] = true
				}
			}

			if count > result {
				result = count
			}
		}
	}

	return result
}
