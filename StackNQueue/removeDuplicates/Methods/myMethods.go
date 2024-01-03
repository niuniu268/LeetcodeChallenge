package Methods

func RemovedDuplicates(s string) string {
	hash := []byte(s)
	stack := make([]byte, 0)

	if hash == nil {

		return ""

	}
	for i := range hash {
		if i == 0 {
			stack = append(stack, hash[i])
		} else {
			if hash[i] == hash[i-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, hash[i])
			}
		}
	}

	return string(stack)
}
