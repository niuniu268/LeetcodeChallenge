package main

import (
	Methods2 "Leetcode/Array/SortedSquares/Methods"
	"fmt"
)

func main() {

	nums := []int{
		3, -2, 1, 5,
	}

	Nnums1 := Methods2.Mymethod(nums)
	Nnums2 := Methods2.TwoPoints(nums)

	fmt.Printf("old array: %v, new array: %v \n", nums, Nnums1)
	fmt.Printf("old array: %v, new array: %v \n", nums, Nnums2)

}
