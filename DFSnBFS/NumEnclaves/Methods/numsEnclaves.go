package methods

var count int
var dir = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func dfs(grid [][]int, x, y int) {
	if grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0
	count++

	for i := 0; i < 4; i++ {
		nextX := x + dir[i][0]
		nextY := y + dir[i][1]

		if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
			continue
		}
		dfs(grid, nextX, nextY)
	}
}

func NumEnclaves(grid [][]int) int {
	count = 0

	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
			dfs(grid, i, 0)
		}
		if grid[i][len(grid[0])-1] == 1 {
			dfs(grid, i, len(grid[0])-1)
		}
	}

	for j := 1; j < len(grid[0])-1; j++ {
		if grid[0][j] == 1 {
			dfs(grid, 0, j)
		}
		if grid[len(grid)-1][j] == 1 {
			dfs(grid, len(grid)-1, j)
		}
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j)
			}
		}
	}
	return count
}
