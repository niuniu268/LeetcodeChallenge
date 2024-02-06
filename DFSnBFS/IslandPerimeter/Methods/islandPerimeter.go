package Methods

func islandPerimeter1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += 4
				// 上下左右四个方向
				if i > 0 && grid[i-1][j] == 1 {
					res--
				} // 上边有岛屿
				if i < m-1 && grid[i+1][j] == 1 {
					res--
				} // 下边有岛屿
				if j > 0 && grid[i][j-1] == 1 {
					res--
				} // 左边有岛屿
				if j < n-1 && grid[i][j+1] == 1 {
					res--
				} // 右边有岛屿
			}
		}
	}
	return res
}
