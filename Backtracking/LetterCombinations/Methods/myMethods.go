package Methods

import (
	"strconv"
)

var (
	m      []string
	path   []byte
	result []string
)

func letterCombinations(digits string) []string {

	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, result = make([]byte, 0), make([]string, 0)

	var combine func(digits string)
	combine = func(digits string) {
		if len(digits) == 0 {
			w := string(path)

			result = append(result, w)

			return
		}

		parseInt, err := strconv.ParseInt(digits, 10, 64)

		if err != nil && digits != "" {
			s := m[parseInt%10-2]
			digits = strconv.Itoa(int(parseInt / 10))

			for i := 0; i < len(s); i++ {
				path = append(path, s[i])
				combine(digits)
				path = path[:len(path)-1]

			}
		}

	}

	combine(digits)

	return result
}
