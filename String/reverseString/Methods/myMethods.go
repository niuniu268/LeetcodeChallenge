package Methods

func reverseString(arr []byte) []byte {

	left := 0
	right := len(arr) - 1

	for left < right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp
		left++
		right--
	}

	return arr

}
