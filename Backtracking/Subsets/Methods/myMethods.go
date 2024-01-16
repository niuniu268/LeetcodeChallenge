package Methods

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	startIndex := 0

	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		tmp := make([]int, len(path))
		copy(tmp, path)

		if startIndex >= len(nums) {
			result = append(result, tmp)
			return
		}

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(nums, startIndex+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, startIndex)
	return result

}
