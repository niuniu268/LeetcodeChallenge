package main

import (
	"Leetcode/HashTable/Intersection/Methods"
	"fmt"
)

func main() {
	arr1 := []int{2, 3, 2}
	arr2 := []int{3, 2, 7, 4, 6}
	intersection := Methods.Intersection(arr1, arr2)

	answer := Methods.InterSectionAnswer(arr1, arr2)

	fmt.Printf("intersection arry: %v\n", intersection)
	fmt.Printf("intersection arry: %v\n", answer)
}
