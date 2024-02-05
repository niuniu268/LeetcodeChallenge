package Methods

func findLen(nums1, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)

	for index := range dp {
		dp[index] = make([]int, len(nums2)+1)
	}

	dp[0][0] = 0
	result := 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			if dp[i+1][j+1] > result {
				result = dp[i+1][j+1]
			}
		}
	}

	return result
}
