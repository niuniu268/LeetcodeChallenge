package Methods

func lenOfLIS(nums []int) int {

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
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] > result {

			result = dp[i]

		}
	}

	return result
}
