package Methods

func dailyTemp(nums []int) []int {
	results := make([]int, len(nums))

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				results[i] = j - i
				break
			} else {
				results[i] = 0
			}
		}
	}
	// Add 0 for the last day
	results = append(results, 0)

	return results
}
