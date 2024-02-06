package Methods

var result = 0
var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(grid [][]int, x, y int) {
	if grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0

	result++

	for i := 0; i < 4; i++ {
		nextX := x + direction[i][0]
		nextY := y + direction[i][1]

		if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid) {
			continue
		}

		dfs(grid, nextX, nextY)

	}
}
func islandPerimeter(grid [][]int) int {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 1 {
				dfs(grid, i, j)
			}

		}
	}

	return result
}
