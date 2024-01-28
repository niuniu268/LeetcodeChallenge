package Methods

func findTargetSumWays2D(nums []int, target int) int {

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if abs(target) > sum || (target+sum)%2 == 1 {
		return 0
	}

	bagSize := (target + sum) / 2

	dp := make([][]int, len(nums))

	for i := range dp {
		dp[i] = make([]int, bagSize+1)
		for j := range dp[i] {
			dp[i][j] = 0
		}
		dp[i][0] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= bagSize; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
		}

	}
	return dp[len(nums)-1][bagSize]
}
