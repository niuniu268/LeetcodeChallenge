package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func constructMaximumBinaryTree(nums []int) *Methods.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 找到最大值
	index := findMax(nums)
	// 构造二叉树
	root := &Methods.TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),   //左半边
		Right: constructMaximumBinaryTree(nums[index+1:]), //右半边
	}
	return root
}
func findMax(nums []int) (index int) {
	for i, v := range nums {
		if nums[index] < v {
			index = i
		}
	}
	return
}
