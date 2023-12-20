package Methods

func TwoSum(arr []int, target int) []int {
	map1 := make(map[int]int)
	result := make([]int, 0)

	for i, v := range arr {
		for i2, v2 := range map1 {
			if target-i2 == v {
				result = append(result, i, v2)
			}

		}
		map1[v] = i

	}

	return result
}
