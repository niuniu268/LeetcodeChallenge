package Methods

func IsAnagram(s, t string) bool {

	c := make([]int, 26)
	for i := range s {
		c[s[i]-'a'] += 1
	}
	for i := range t {
		c[t[i]-'a'] -= 1

	}

	for i := range c {
		if c[i] != 0 {
			return false
		}
	}

	return true
}
