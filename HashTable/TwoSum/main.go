package main

import (
	"Leetcode/HashTable/TwoSum/Methods"
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5}

	target := 5
	sum := Methods.TwoSum(arr, target)

	fmt.Printf("array: %v \n", sum)
}
