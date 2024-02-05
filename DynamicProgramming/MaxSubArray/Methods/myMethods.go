package Methods

func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	result := 0
	dp[0] = nums[0]

	for i := 0; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], dp[i])
		if dp[i] > result {
			result = dp[i]
		}

	}
	return result
}
