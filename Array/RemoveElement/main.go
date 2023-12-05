package main

import (
	Methods2 "Leetcode/Array/RemoveElement/Methods"
	"fmt"
)

func main() {

	nums := []int{
		1, 2, 3, 4, 5, 5, 6, 7, 8, 4, 3, 2,
	}
	target := 5

	var lens, arr = Methods2.MyMethod(nums, target)
	lens1, arr1 := Methods2.TwoPoints(nums, target)

	fmt.Printf("lens %v, arr %v \n", lens, arr)
	fmt.Printf("lens %v, arr %v \n ", lens1, arr1)

}
