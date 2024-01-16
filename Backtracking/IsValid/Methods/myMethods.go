package Methods

import (
	"strconv"
	"strings"
)

func isValid(num string) []string {
	result := make([]string, 0)
	path := make([]string, 0)
	startIndex := 0

	var backtrack func(num string, startIndex int)

	backtrack = func(num string, startIndex int) {
		if startIndex == len(num) && len(path) == 4 {
			for _, s := range path {
				if s[0] == '0' && len(s) > 1 {
					return
				}
				value, err := strconv.Atoi(s)
				if err != nil {
					return
				}
				if value > 255 {
					return
				}
			}
			result = append(result, strings.Join(path, "."))
			return
		}

		for i := startIndex; i < len(num); i++ {
			substr := num[startIndex : i+1]
			path = append(path, substr)
			backtrack(num, startIndex+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(num, startIndex)

	return result

}
