package Methods

func palindrome(input string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	startIndex := 0

	var backtracking func(input string, startIndex int)

	backtracking = func(input string, startIndex int) {
		if startIndex >= len(input) {
			result = append(result, path)
			return
		}

		for i := startIndex; i < len(input); i++ {
			if isPalindrome(input, startIndex, i) {

				str := input[startIndex : i+1]
				path = append(path, str)
				backtracking(input, startIndex+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtracking(input, startIndex)

	return result
}

func isPalindrome(input string, start, end int) bool {

	for start < end {
		if input[start] != input[end] {
			return false
		}
		start++
		end--
	}
	return true
}
