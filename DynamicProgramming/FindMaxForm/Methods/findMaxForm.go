package Methods

func findMaxForm(strs []string, m int, n int) int {
	// 定义数组
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1 个数
		//或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum
		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
		//fmt.Println(dp)
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//输入：strs = ["1010", "1100","10","1", "0"], m = 2, n = 2
//The result you would get from running the provided findMaxForm function is 3.
//
//Explanation: The maximum size of a subset with at most m=2 zeros and n=2 ones can be achieved by selecting the strings "10", "1", and "0". This results in a total of 3 strings while satisfying the zero (2) and one (2) constraints.
