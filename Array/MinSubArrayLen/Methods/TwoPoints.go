package Methods

func TwoPoints(nums []int, target int) int {

	i := 0
	l := len(nums)
	sum := 0
	result := l + 1

	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			sublengthen := j - i + 1
			if sublengthen < result {
				result = sublengthen
			}
			sum -= nums[i]
			i++
		}

	}

	if result == l+1 {
		return 0
	} else {
		return result
	}
}
