package Methods

import (
	"slices"
)

func RepeatedSub(word string) bool {
	arr := []byte(word)
	for i := 1; i <= len(arr)/2; i++ {
		subarr := arr[:i]

		if len(arr)%len(subarr) == 0 {
			for j := 1; j < len(arr)/len(subarr); j++ {
				start := j * len(subarr)
				end := (j + 1) * len(subarr)
				if !slices.Equal(arr[start:end], subarr) {
					break
				}
				if j == len(arr)/len(subarr)-1 {
					return true
				}
			}
		}
	}
	return false
}
