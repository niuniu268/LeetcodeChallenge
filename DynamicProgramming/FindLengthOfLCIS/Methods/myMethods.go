package Methods

func findLenOfLCIS(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	result := 0

	dp := make([]int, len(nums))

	// 初始化，所有的元素都应该初始化为1
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {

		if dp[i] > dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}

		if dp[i] > result {

			result = dp[i]

		}

	}

	return result
}
