package Methods

func combine(n, k int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(n, k, startIndex int)
	backtracking = func(n, k, startIndex int) {
		if len(path) == k {
			result = append(result, path)

			return
		}

		for i := startIndex; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)

	return result
}

func combine2(target, k int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(target, k, startIndex int)
	backtracking = func(target, k, startIndex int) {
		if len(path) == k {
			sum := 0
			for _, v := range path {
				sum += v
			}
			if sum == target {
				result = append(result, path)
			}

			return
		}

		for i := startIndex; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking(9, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(target, k, 1)

	return result
}
