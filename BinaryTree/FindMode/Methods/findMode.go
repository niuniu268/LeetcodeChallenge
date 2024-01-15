package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func findMode(root *Methods.TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *Methods.TreeNode
	var travel func(node *Methods.TreeNode)
	travel = func(node *Methods.TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}
