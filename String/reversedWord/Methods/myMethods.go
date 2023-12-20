package Methods

import (
	"slices"
	"strings"
)

func ReversedWord(phrase string) string {
	phraseArr := strings.Split(phrase, " ")

	left := 0
	right := len(phraseArr) - 1

	for left < right {
		phraseArr[left], phraseArr[right] = phraseArr[right], phraseArr[left]
		left++
		right--
	}

	for i, s := range phraseArr {
		if s == " " {
			slices.Delete(phraseArr, i, i+1)
		}
	}

	return strings.Join(phraseArr, " ")
}
