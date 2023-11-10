package main

import "fmt"

func main() {

	arr := make([]int, 5)
	arr = append(arr, 1, 2, 3, 4, 5)

	var arr2 = [4][3]int{
		{1, 2, 3},
		{3, 4, 5},
		{6, 7, 8},
		{9, 9, 9}}

	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])
	fmt.Println(&arr[3])
	fmt.Println(&arr[4])

	fmt.Println(&arr2[1][1])
	fmt.Println(&arr2[1][2])

	fmt.Println("============")

	var nums = []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	fmt.Println(search(nums, 6))

}

func search(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {

		if i == target {

			return i
		}
	}

	return -1

}
