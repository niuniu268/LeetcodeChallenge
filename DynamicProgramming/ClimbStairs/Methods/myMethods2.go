package Methods

func climbStairsDp(step int, target int) int {
	nums := make([]int, step)
	dp := make([]int, target+1)

	for i := 0; i < step; i++ {
		nums[i] = i + 1
	}
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j < target+1; j++ {
			if j > nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}
