package Methods

func intBreak(num int) int {

	dp := make([]int, num)
	dp[1] = 1
	dp[0] = 1

	for i := 2; i < num; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))

		}

	}
	return dp[num-1]
}
