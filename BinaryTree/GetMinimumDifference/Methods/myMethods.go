package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func GetMinimun(root *Methods.TreeNode) int {
	result := 0

	var pre *Methods.TreeNode

	if root == nil {
		return result
	}

	var traversal func(root *Methods.TreeNode)

	traversal = func(root *Methods.TreeNode) {

		if root == nil {
			return
		}
		traversal(root.Left)

		if pre != nil && result > root.Val-pre.Val {
			result = root.Val - pre.Val
		}
		pre = root
		traversal(root.Right)

	}

	traversal(root)

	return result

}
