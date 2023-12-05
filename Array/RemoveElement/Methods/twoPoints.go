package Methods

func TwoPoints(nums []int, val int) (int, []int) {

	left := 0
	right := len(nums) - 1

	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}

	}

	return left, nums
}
