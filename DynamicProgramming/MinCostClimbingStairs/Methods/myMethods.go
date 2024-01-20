package Methods

func minClimb(cost []int) int {

	if len(cost) <= 2 {
		return cost[len(cost)-1]
	}

	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}

	return min(dp[len(dp)-2]+cost[len(cost)-2], dp[len(dp)-1])
}
