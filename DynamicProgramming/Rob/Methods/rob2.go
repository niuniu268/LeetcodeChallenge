package Methods

// 打家劫舍Ⅱ 动态规划
// 时间复杂度O(n) 空间复杂度O(n)
func rob(nums []int) int {
	// 如果长度为0或1，那么有没有环的限制都一样
	if len(nums) <= 1 {
		return robWithoutCircle(nums)
	}

	// 否则，去头或去尾，取最大
	res1 := robWithoutCircle(nums[:len(nums)-1])
	res2 := robWithoutCircle(nums[1:])

	return max(res1, res2)
}

// 原始的打家劫舍版
func robWithoutCircle(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
