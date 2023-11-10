package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 3, 4, 5,
	}
	target := 4

	position := Bisection(nums, target)
	fmt.Println(position)

}

func Bisection(nums []int, target int) int {

	high := len(nums) - 1
	low := 0

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1

}
