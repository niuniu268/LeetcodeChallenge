package Methods

func maxProfit(prices []int) int {
	if len(prices) == 0 {

		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(-prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
	}

	return dp[len(prices)-1][1]
}
