package Methods

func ReverseString(word string, k int) string {

	groupNum := len(word) / (2 * k)
	lastGroupNum := len(word) % (2 * k)

	arr := []byte(word)

	for i := 0; i < groupNum; i++ {
		left := i * 2 * k
		right := i*2*k + k - 1
		for left < right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			left++
			right--
		}
	}

	if lastGroupNum < k {
		left := groupNum * 2 * k
		right := len(word) - 1
		for left < right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			left++
			right--
		}
	} else {
		left := groupNum * 2 * k
		right := groupNum*2*k + k - 1
		for left < right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			left++
			right--
		}
	}

	return string(arr[:])
}
