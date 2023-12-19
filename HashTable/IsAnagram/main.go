package main

import (
	"Leetcode/HashTable/IsAnagram/Methods"
	"fmt"
)

func main() {

	test := Methods.IsAnagram("abcc", "cba")

	answer := Methods.IsAnagramAnswer("abc", "ccb")

	fmt.Printf("test result: %v\n", test)
	fmt.Printf("test result: %v\n", answer)
}
