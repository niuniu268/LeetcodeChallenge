package Methods

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	mapCheck := make(map[int]bool, len(nums))

	var backtracking func(nums []int, mapCheck map[int]bool)

	backtracking = func(nums []int, mapCheck map[int]bool) {

		if len(path) == len(nums) {
			result = append(result, path)
			return
		}

		for i := 0; i < len(nums); i++ {
			if mapCheck[i] {
				break
			}
			path = append(path, nums[i])
			mapCheck[i] = true
			backtracking(nums, mapCheck)
			path = path[:len(path)-1]
			mapCheck[i] = false
		}
	}

	backtracking(nums, mapCheck)

	return result
}
