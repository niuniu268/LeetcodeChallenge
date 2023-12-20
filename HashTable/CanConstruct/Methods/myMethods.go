package Methods

func CanConstruct(word1 string, word2 string) bool {
	ints := make([]int, 26)

	for _, v := range word1 {
		ints[v-'a']++
	}

	for _, v := range word2 {
		ints[v-'a']--
	}

	for _, v := range ints {
		if v > 0 {
			return false
		}
	}

	return true

}
