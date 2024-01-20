package Methods

//第二种思路: dp[i]表示从i层起跳所需要支付的最小费用
//递推公式:
//i<n :dp[i] = min(dp[i-1],dp[i-2])+cost[i]
//i==n:dp[i] = min(dp[i-1],dp[i-2]) (登顶)

func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i <= n; i++ {
		if i < n {
			dp[i] = min2(dp[i-1], dp[i-2]) + cost[i]
		} else {
			dp[i] = min2(dp[i-1], dp[i-2])
		}
	}
	return dp[n]
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
