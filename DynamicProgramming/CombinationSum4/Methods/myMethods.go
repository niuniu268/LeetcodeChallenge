package Methods

func combinedSum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i < target+1; i++ {
		for j := 0; j < len(nums)+1; j++ {
			if i > nums[j] {
				dp[i] = dp[i-nums[j]]
			}

		}
	}
	return dp[target]
}
