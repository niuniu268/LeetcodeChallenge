package Methods

func MyMethod(nums []int, target int) (int, []int) {

	lens := len(nums)
	var count = 0

	for i := lens - 1; i >= 0; i-- {
		if nums[i] == target {
			count++
			var tmp = nums[lens-count]
			nums[i] = tmp
			nums[lens-count] = target

		}

	}

	return lens - count, nums

}
