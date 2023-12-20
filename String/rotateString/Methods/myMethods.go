package Methods

func RotateString(wordString string, num int) string {

	slice := []byte(wordString)

	slice = reverse(slice)
	slice1 := reverse(slice[:num])
	slice2 := reverse(slice[num:])
	slice1 = append(slice1, slice2...)

	return string(slice1[:])
}

func reverse(slice []byte) []byte {

	left := 0
	right := len(slice) - 1

	for left < right {
		slice[right], slice[left] = slice[left], slice[right]
		left++
		right--
	}

	return slice
}
