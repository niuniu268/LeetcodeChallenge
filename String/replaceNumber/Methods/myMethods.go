package Methods

import (
	"unicode"
)

func ReplaceNumber(str string) string {
	arr := []byte(str)

	count := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(arr[i])) {
			count++
		}
	}

	point2 := len(arr) + count*5 - 1
	arr2 := make([]byte, point2+1)
	for point1 := len(arr) - 1; point1 >= 0; point1-- {
		if !unicode.IsNumber(rune(arr[point1])) {
			arr2[point2] = arr[point1]
			point2--
		} else {
			arr2[point2] = 'r'
			point2--
			arr2[point2] = 'e'
			point2--
			arr2[point2] = 'b'
			point2--
			arr2[point2] = 'm'
			point2--
			arr2[point2] = 'u'
			point2--
			arr2[point2] = 'n'
			point2--
		}
	}

	return string(arr2[:])
}
