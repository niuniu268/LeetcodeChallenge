package Methods

func numDistinct(nums1, nums2 string) int {
	dp := make([][]int, len(nums2)+1)
	for index := range dp {
		dp[index] = make([]int, len(nums1)+1)
		dp[index][0] = 1

	}

	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]

			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums1)][len(nums2)]
}
