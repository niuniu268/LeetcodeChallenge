package Methods

var (
	res []string
)

func letterCombinationsAnswer(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res
	}
	dfs(digits, 0)
	return res
}
func dfs(digits string, start int) {
	if len(path) == len(digits) { //终止条件，字符串长度等于digits的长度
		tmp := string(path)
		res = append(res, tmp)
		return
	}
	digit := int(digits[start] - '0') // 将index指向的数字转为int（确定下一个数字）
	str := m[digit-2]                 // 取数字对应的字符集（注意和map中的对应）
	for j := 0; j < len(str); j++ {
		path = append(path, str[j])
		dfs(digits, start+1)
		path = path[:len(path)-1]
	}
}
