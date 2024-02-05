package Methods

func isSubsequence1D(s string, t string) bool {
	dp := make([]int, len(s)+1)
	for i := 1; i <= len(t); i++ {
		for j := len(s); j >= 1; j-- {
			if t[i-1] == s[j-1] {
				dp[j] = dp[j-1] + 1
			}
		}
	}
	return dp[len(s)] == len(s)
}
