package Methods

func test_1_wei_bag_problem(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([]int, bagWeight+1)
	// 递推顺序
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test_1_wei_bag_problem(weight, value, 4)
}
