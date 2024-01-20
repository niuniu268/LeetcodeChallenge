package Methods

func dpFib(num int) int {
	dp := make([]int, 0)
	if num <= 1 {
		return num
	}
	dp[0] = 0
	dp[1] = 1

	for i := 2; i < num; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[num]
}
