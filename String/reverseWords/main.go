package main

import (
	"Leetcode/String/reverseWords/Methods"
	"fmt"
)

func main() {
	reverseString := Methods.ReverseString("abcdefg", 2)
	fmt.Printf("input word: %v, and reversed word: %v\n", "abcdefg", reverseString)

}
