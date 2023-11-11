package main

import (
	"Leetcode/RemoveElement/Methods"
	"fmt"
)

func main() {

	nums := []int{
		1, 2, 3, 4, 5, 5, 6, 7, 8, 4, 3, 2,
	}
	target := 5

	var lens, arr = Methods.MyMethod(nums, target)
	lens1, arr1 := Methods.TwoPoints(nums, target)

	fmt.Printf("lens %v, arr %v \n", lens, arr)
	fmt.Printf("lens %v, arr %v \n ", lens1, arr1)

}
