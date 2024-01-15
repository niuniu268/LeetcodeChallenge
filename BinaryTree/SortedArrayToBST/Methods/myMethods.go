package Methods

import (
	`Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods`
)

func sortArrayToBEST(arr []int) *Methods.TreeNode {
	left := 0
	right := len(arr) - 1

	var traversal func(arr []int, left, right int) *Methods.TreeNode

	traversal = func(arr []int, left, right int) *Methods.TreeNode {
		if left > right {

			return nil
		}

		mid := left + (right-left)/2

		node := &Methods.TreeNode{Val: arr[mid]}

		node.Left = traversal(arr, left, mid-1)
		node.Right = traversal(arr, mid+1, right)

		return node
	}

	return traversal(arr, left, right)
}
