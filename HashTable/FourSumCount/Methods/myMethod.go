package Methods

func FourSumCount(arr1 []int, arr2 []int, arr3 []int, arr4 []int) int {

	result := 0
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			for _, v3 := range arr3 {
				for _, v4 := range arr4 {
					if v1+v2+v3+v4 == 0 {
						result++
					}
				}
			}
		}
	}

	return result
}
