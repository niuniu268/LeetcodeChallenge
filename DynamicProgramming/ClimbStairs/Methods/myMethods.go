package Methods

func climb(stairs int) int {
	if stairs <= 1 {
		return stairs
	}

	dp := make([]int, stairs)
	dp[0], dp[1] = 1, 2

	for i := 3; i <= stairs; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[stairs-1]
}
